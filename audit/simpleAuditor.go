package audit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codenotary/immuproof/meta"
	"github.com/codenotary/immuproof/status"
	"github.com/vchain-us/ledger-compliance-go/grpcclient"
	"github.com/vchain-us/ledger-compliance-go/schema"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"log"
	"strings"
	"time"
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

func (a *simpleAuditor) Audit() error {
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
	go func() {
		for {
			var i int
			for i = 0; i < len(a.apiKeys); i++ {
				ak := a.apiKeys[i]

				signerID, err := GetSignerIDFromApiKey(ak)
				if err != nil {
					log.Printf("failed to get signer id from api key: %s", ak)
					continue
				}

				select {
				case <-a.done:
					return
				case <-ticker.C:
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
						statusReport.Status = err.Error()
						a.statusMap.Add(statusReport)
						log.Printf("error checking consistency: %v", err)
					} else {
						statusReport.Status = status.Status_NORMAL
						statusReport.NewStateHash = cResp.NewStateHash
						statusReport.PrevStateHash = cResp.PrevStateHash
						a.statusMap.Add(statusReport)
					}
					err = a.SaveStatusMap()
					if err != nil {
						log.Printf("failed to save status map: %v", err)
					}
				}
			}
			i = 0
		}
	}()

	<-a.done
	return nil
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
	err = ioutil.WriteFile(meta.StateMapFileName, j, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (a *simpleAuditor) LoadStatusMap() error {
	j, err := ioutil.ReadFile(meta.StateMapFileName)
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
