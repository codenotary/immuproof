package audit

import (
	"context"
	"errors"
	"github.com/codenotary/immuproof/status"
	"github.com/vchain-us/ledger-compliance-go/grpcclient"
	"log"
	"strings"
	"time"
)

type simpleAuditor struct {
	done      chan struct{}
	client    grpcclient.LcClientIf
	statusMap *status.StatusReportMap
}

func NewSimpleAuditor(client grpcclient.LcClientIf, statusMap *status.StatusReportMap) *simpleAuditor {
	return &simpleAuditor{client: client, statusMap: statusMap, done: make(chan struct{})}
}

func (a *simpleAuditor) Audit() error {
	ticker := time.NewTicker(500 * time.Millisecond)
	signerID, err := GetSignerIDFromApiKey(a.client.(*grpcclient.LcClient).ApiKey)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case <-a.done:
				return
			case <-ticker.C:
				untrustedState, err := a.client.CurrentState(context.TODO())
				if err != nil {
					log.Printf("error getting untrusted state: %v", err)
				}
				err = a.client.ConsistencyCheck(context.TODO(), untrustedState)
				if err != nil {
					if strings.Contains(err.Error(), "corrupted data") {
						statusReport := status.StatusReport{
							SignerID: signerID,
							Status:   status.Status_CORRUPTED_DATA,
							Time:     time.Now(),
						}
						a.statusMap.Add(statusReport)
					}
					statusReport := status.StatusReport{
						SignerID: signerID,
						Status:   err.Error(),
						Time:     time.Now(),
					}
					a.statusMap.Add(statusReport)
					log.Printf("error checking consistency: %v", err)
				} else {
					statusReport := status.StatusReport{
						SignerID: signerID,
						Status:   status.Status_NORMAL,
						Time:     time.Now(),
					}
					a.statusMap.Add(statusReport)
				}
			}
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
