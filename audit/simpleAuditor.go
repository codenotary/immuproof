package audit

import (
	"context"
	"errors"
	"github.com/codenotary/immuproof/status"
	"github.com/vchain-us/ledger-compliance-go/grpcclient"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
	"time"
)

type simpleAuditor struct {
	done      chan struct{}
	client    grpcclient.LcClientIf
	statusMap *status.StatusReportMap
	apiKeys   []string
}

func NewSimpleAuditor(client grpcclient.LcClientIf, statusMap *status.StatusReportMap) *simpleAuditor {
	if client == nil || client.IsConnected() == false {
		log.Fatal("client is nil or not connected")
	}
	return &simpleAuditor{client: client, statusMap: statusMap, done: make(chan struct{}), apiKeys: []string{}}
}

func (a *simpleAuditor) AddApiKey(apiKey string) {
	a.apiKeys = append(a.apiKeys, apiKey)
}

func (a *simpleAuditor) Audit() error {
	ticker := time.NewTicker(1000 * time.Millisecond)
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
					err = a.client.ConsistencyCheck(ctx)
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
						a.statusMap.Add(statusReport)
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
