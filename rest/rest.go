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

package rest

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"

	"github.com/codenotary/immuproof/status"
)

//go:embed internal/embed
var content embed.FS

type restServer struct {
	port          string
	webCertFile   string
	webKeyFile    string
	statusHandler *statusHandler
	countHandler  *countHandler
}

type statusHandler struct {
	statusMap *status.StatusReportMap
}

type countHandler struct {
	statusMap *status.StatusReportMap
}

func NewRestServer(statusMap *status.StatusReportMap, port, webCertFile, webKeyFile string) *restServer {
	return &restServer{
		port:        port,
		webCertFile: webCertFile,
		webKeyFile:  webKeyFile,
		statusHandler: &statusHandler{
			statusMap: statusMap,
		},
		countHandler: &countHandler{
			statusMap: statusMap,
		},
	}
}

func (s *restServer) Serve() error {
	log.Printf("Starting REST server on port %s", s.port)
	log.Print("UI is exposed on /")
	log.Print("REST status history is exposed on /api/status")
	log.Print("REST notarization's counter are exposed on /api/notarization/count")

	mux := http.NewServeMux()

	muxCors := cors.Default().Handler(mux)

	index, err := fs.Sub(content, "internal/embed")
	if err != nil {
		return err
	}

	mux.Handle("/", http.FileServer(http.FS(index)))
	mux.Handle("/api/status", s.statusHandler)
	mux.Handle("/api/notarization/count", s.countHandler)

	if s.webCertFile != "" && s.webKeyFile != "" {
		log.Print("REST server is using HTTPS")
		return http.ListenAndServeTLS(fmt.Sprintf(":%s", s.port), s.webCertFile, s.webKeyFile, muxCors)
	} else {
		log.Print("REST server is using HTTP")
		return http.ListenAndServe(fmt.Sprintf(":%s", s.port), muxCors)
	}
}

func (s *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.statusMap.GetAllByLedger())
}

type ledgerCounter struct {
	NewNotarizationsCount uint64    `json:"newNotarizationsCount"`
	CollectTime           time.Time `json:"collectTime"`
	CollectTimeZone       string    `json:"collectTimeZone"`
}

func (s *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reports := s.statusMap.GetAllByLedger()
	counts := make(map[string][]*ledgerCounter)
	for id, checks := range reports {
		lcs := make([]*ledgerCounter, 0)
		var old *status.StatusReport
		for _, check := range checks {
			nc := uint64(0)
			if old != nil {
				nc = check.NewTxID - old.NewTxID
				if old.NewTxID > check.NewTxID {
					nc = 0
				}
			}
			lc := &ledgerCounter{
				NewNotarizationsCount: nc,
				CollectTime:           check.Time,
				CollectTimeZone:       check.TimeZone,
			}
			lcs = append(lcs, lc)
			old = check
		}
		counts[id] = lcs
	}
	json.NewEncoder(w).Encode(counts)
}
