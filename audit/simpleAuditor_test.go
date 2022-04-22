/*
Copyright © 2022 CodeNotary, Inc. All rights reserved
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package audit

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/codenotary/immuproof/cnc/cnctest"

	"github.com/codenotary/immuproof/status"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	sdk "github.com/vchain-us/ledger-compliance-go/grpcclient"
	"github.com/vchain-us/ledger-compliance-go/schema"
)

func TestSimpleAuditor(t *testing.T) {

	viper.Set("audit-interval", "1s")
	viper.Set("state-history-file", "tmpStateCache.json")
	historySize := 2
	defer os.Remove("tmpStateCache.json")

	aks := []string{"signerID1.ak1", "signerID2.ak2"}
	cMock := &cnctest.LcClientMock{
		ConsistencyCheckF: func(ctx context.Context) (*sdk.ConsistencyCheckResponse, error) {
			return &sdk.ConsistencyCheckResponse{
				PrevTxID:      0,
				PrevStateHash: "",
				NewTxID:       0,
				NewStateHash:  "",
			}, nil
		},
		IsConnectedF: func() bool {
			return true
		},
		FeatsF: func(ctx context.Context) (*schema.Features, error) {
			return &schema.Features{
				Feat:      []string{"immuproof"},
				Version:   "",
				BuildTime: "",
				Component: "",
			}, nil
		},
	}
	statusReportMap := status.NewStatusReportMap(historySize)
	simpleAuditor := NewSimpleAuditor(cMock, statusReportMap, viper.GetDuration("audit-interval"))
	for _, a := range aks {
		simpleAuditor.AddApiKey(a)
	}
	go func() {
		err := simpleAuditor.Start()
		require.NoError(t, err)
	}()

	time.Sleep(time.Second * 4)
	simpleAuditor.Stop()

	tmpStatusReportMap := status.NewStatusReportMap(historySize)
	j, err := ioutil.ReadFile("tmpStateCache.json")
	require.NoError(t, err)
	err = json.Unmarshal(j, &tmpStatusReportMap)
	require.NoError(t, err)
	require.True(t, len(tmpStatusReportMap.M) == historySize)
}

func TestSimpleAuditorCorruptedData(t *testing.T) {
	type testData struct {
		errorMsg      string
		expectedError string
	}

	tds := []testData{
		{
			errorMsg:      "corrupted data",
			expectedError: status.Status_CORRUPTED_DATA,
		},
		{
			errorMsg:      "data is corrupted",
			expectedError: status.Status_CORRUPTED_DATA,
		},
		{
			errorMsg:      "other error",
			expectedError: status.Status_UNKNOWN,
		},
	}

	defer os.Remove("tmpStateCache.json")

	for _, td := range tds {
		os.Remove("tmpStateCache.json")
		viper.Set("audit-interval", "1s")
		viper.Set("state-history-file", "tmpStateCache.json")
		historySize := 2

		aks := []string{"signerID1.ak1", "signerID2.ak2"}
		cMock := &cnctest.LcClientMock{
			ConsistencyCheckF: func(ctx context.Context) (*sdk.ConsistencyCheckResponse, error) {
				return nil, fmt.Errorf(td.errorMsg)
			},
			IsConnectedF: func() bool {
				return true
			},
			FeatsF: func(ctx context.Context) (*schema.Features, error) {
				return &schema.Features{
					Feat: []string{"immuproof"},
				}, nil
			},
		}
		statusReportMap := status.NewStatusReportMap(historySize)
		simpleAuditor := NewSimpleAuditor(cMock, statusReportMap, viper.GetDuration("audit-interval"))
		for _, a := range aks {
			simpleAuditor.AddApiKey(a)
		}
		go func() {
			err := simpleAuditor.Start()
			require.NoError(t, err)
		}()

		time.Sleep(time.Second * 4)
		simpleAuditor.Stop()

		tmpStatusReportMap := status.NewStatusReportMap(historySize)
		j, err := ioutil.ReadFile("tmpStateCache.json")
		require.NoError(t, err)
		err = json.Unmarshal(j, &tmpStatusReportMap)
		require.NoError(t, err)
		require.True(t, len(tmpStatusReportMap.M) == historySize)
		allByLed := tmpStatusReportMap.GetAllByLedger()
		for _, led := range allByLed {
			for _, s := range led {
				require.True(t, s.Status == td.expectedError)
				break
			}
		}
	}
}

func TestSimpleAuditorFeatNotSupported(t *testing.T) {

	viper.Set("audit-interval", "1s")
	viper.Set("state-history-file", "tmpStateCache.json")
	historySize := 2
	defer os.Remove("tmpStateCache.json")

	aks := []string{"signerID1.ak1", "signerID2.ak2"}
	cMock := &cnctest.LcClientMock{
		IsConnectedF: func() bool {
			return true
		},
		FeatsF: func(ctx context.Context) (*schema.Features, error) {
			return &schema.Features{
				Feat: []string{},
			}, nil
		},
	}
	statusReportMap := status.NewStatusReportMap(historySize)
	simpleAuditor := NewSimpleAuditor(cMock, statusReportMap, viper.GetDuration("audit-interval"))
	for _, a := range aks {
		simpleAuditor.AddApiKey(a)
	}

	err := simpleAuditor.Start()
	require.Contains(t, err.Error(), "doesn't support immuproof feature")
}
