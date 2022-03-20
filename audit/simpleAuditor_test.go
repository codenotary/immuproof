package audit

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/codenotary/immuproof/cnc"
	"github.com/codenotary/immuproof/status"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	sdk "github.com/vchain-us/ledger-compliance-go/grpcclient"
	"github.com/vchain-us/ledger-compliance-go/schema"
)

func TestSimpleAuditor(t *testing.T) {

	viper.Set("audit-interval", "1s")
	viper.Set("state-cache-file", "tmpStateCache.json")
	viper.Set("state-cache-size", 2)
	defer os.Remove("tmpStateCache.json")

	aks := []string{"signerID1.ak1", "signerID2.ak2"}
	cMock := &cnc.LcClientMock{
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
	statusReportMap := status.NewStatusReportMap()
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

	tmpStatusReportMap := status.NewStatusReportMap()
	j, err := ioutil.ReadFile("tmpStateCache.json")
	require.NoError(t, err)
	err = json.Unmarshal(j, &tmpStatusReportMap)
	require.NoError(t, err)
	require.True(t, len(tmpStatusReportMap.M) == 2)
}
