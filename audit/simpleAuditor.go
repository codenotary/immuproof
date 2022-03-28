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
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/codenotary/immuproof/status"
	"github.com/spf13/viper"
	"github.com/vchain-us/ledger-compliance-go/grpcclient"
	"github.com/vchain-us/ledger-compliance-go/schema"
	"google.golang.org/grpc/metadata"
)

type simpleAuditor struct {
	done          chan struct{}
	client        grpcclient.LcClientIf
	statusMap     *status.StatusReportMap
	apiKeys       []string
	auditInterval time.Duration
}

func NewSimpleAuditor(client grpcclient.LcClientIf, statusMap *status.StatusReportMap, auditInterval time.Duration) *simpleAuditor {
	if client == nil || client.IsConnected() == false {
		log.Fatal("client is nil or not connected")
	}
	return &simpleAuditor{client: client, statusMap: statusMap, done: make(chan struct{}), apiKeys: []string{}, auditInterval: auditInterval}
}

func (a *simpleAuditor) AddApiKey(apiKey string) {
	a.apiKeys = append(a.apiKeys, apiKey)
}

func (a *simpleAuditor) Start() error {
	f, err := a.client.Feats(context.TODO())
	if err != nil {
		return err
	}
	if _, ok := f.Map()[schema.FeatImmuProof]; !ok {
		return fmt.Errorf("seems that the connected server component `%s` at version `%s` builded at `%s` doesn't support %s feature. Please contact a system administrator", f.Component, f.Version, f.BuildTime, schema.FeatImmuProof)
	}

	err = a.LoadStatusMap()
	if err != nil {
		return err
	}
	ticker := time.NewTicker(a.auditInterval)
	first := true
	for {
		var i int
		for i = 0; i < len(a.apiKeys); i++ {
			ak := a.apiKeys[i]

			signerID, err := GetSignerIDFromApiKey(ak)
			if err != nil {
				log.Printf("failed to get signer id from api key: %s", ak)
				continue
			}
			if first {
				// show something now
				a.collectOne(ak, signerID)
				first = false
			}

			select {
			case <-a.done:
				return nil
			case <-ticker.C:
				a.collectOne(ak, signerID)
			}
		}
		i = 0
	}
}

func (a *simpleAuditor) Stop() {
	a.done <- struct{}{}
}

func GetSignerIDFromApiKey(lcApiKey string) (string, error) {
	ris := strings.Split(lcApiKey, ".")
	if len(ris) > 1 {
		return strings.Join(ris[:len(ris)-1], "."), nil
	}
	return "", errors.New("invalid api key")
}

func (a *simpleAuditor) SaveStatusMap() error {
	j, err := json.Marshal(a.statusMap)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(viper.GetString("state-history-file"), j, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (a *simpleAuditor) LoadStatusMap() error {
	j, err := ioutil.ReadFile(viper.GetString("state-history-file"))
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil
		}
	}
	err = json.Unmarshal(j, &a.statusMap)
	if err != nil {
		return err
	}
	return nil
}

func (a *simpleAuditor) collectOne(ak, signerID string) {
	statusReport := status.StatusReport{
		SignerID: signerID,
		Time:     time.Now(),
	}

	ctx := metadata.AppendToOutgoingContext(context.TODO(), "lc-api-key", ak)
	cResp, err := a.client.ConsistencyCheck(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "corrupted data") {
			statusReport.Status = status.Status_CORRUPTED_DATA
			a.statusMap.Add(statusReport)
		}
		statusReport.Status = status.Status_UNKNOWN
		a.statusMap.Add(statusReport)
		log.Printf("error checking consistency: %v", err)
	} else {
		statusReport.Status = status.Status_NORMAL
		statusReport.NewTxID = cResp.NewTxID
		statusReport.NewStateHash = cResp.NewStateHash
		statusReport.PrevTxID = cResp.PrevTxID
		statusReport.PrevStateHash = cResp.PrevStateHash
		a.statusMap.Add(statusReport)
	}
	err = a.SaveStatusMap()
	if err != nil {
		log.Printf("failed to save status map: %v", err)
	}
}
