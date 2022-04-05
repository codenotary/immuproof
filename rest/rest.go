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
	"path/filepath"
	"text/template"
	"time"

	"github.com/rs/cors"

	"github.com/codenotary/immuproof/status"
)

//go:embed internal/embed
var content embed.FS

type restServer struct {
	address       string
	port          string
	webCertFile   string
	webKeyFile    string
	httpServer    *http.Server
	statusHandler *statusHandler
	countHandler  *countHandler
	webHandler    *webHandler
}

type statusHandler struct {
	statusMap *status.StatusReportMap
}

type countHandler struct {
	statusMap *status.StatusReportMap
}

type webHandler struct {
	address         string
	port            string
	hostedByLogoURL string
	hostedByText    string
	titleText       string
}

func NewRestServer(statusMap *status.StatusReportMap, port, address, webCertFile, webKeyFile, webHostedByLogoURL, webHostedByText, webTitleText string) (*restServer, error) {
	mux := http.NewServeMux()
	muxCors := cors.Default().Handler(mux)

	rs := &restServer{
		port:        port,
		webCertFile: webCertFile,
		webKeyFile:  webKeyFile,
		statusHandler: &statusHandler{
			statusMap: statusMap,
		},
		countHandler: &countHandler{
			statusMap: statusMap,
		},
		webHandler: &webHandler{
			address:         address,
			port:            port,
			hostedByLogoURL: webHostedByLogoURL,
			hostedByText:    webHostedByText,
			titleText:       webTitleText,
		},
		httpServer: &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: muxCors},
	}

	mux.Handle("/", rs.webHandler)
	mux.Handle("/api/status", rs.statusHandler)
	mux.Handle("/api/notarization/count", rs.countHandler)

	return rs, nil
}

func (s *restServer) Serve() error {
	log.Printf("Starting REST server on port %s", s.port)
	log.Print("UI is exposed on /")
	log.Print("REST status history is exposed on /api/status")
	log.Print("REST notarization's counter are exposed on /api/notarization/count")
	if s.webCertFile != "" && s.webKeyFile != "" {
		log.Print("REST server is using HTTPS")
		return s.httpServer.ListenAndServeTLS(s.webCertFile, s.webKeyFile)
	} else {
		log.Print("REST server is using HTTP")
		return s.httpServer.ListenAndServe()
	}
}

func (s *webHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index, err := fs.Sub(content, "internal/embed")
	if err != nil {
		log.Fatalln(err)
	}

	if r.URL.Path == "/js/app.js" || r.URL.Path == "/js/app.js.map" {
		file := filepath.Base(r.URL.Path)
		view := template.Must(template.New("").Delims("{{{", "}}}").ParseFS(index, "js/"+file))

		type env struct {
			PORT, ADDRESS, HOSTED_BY_LOGO_URL, HOSTED_BY_TEXT, TITLE_TEXT string
		}

		e := env{
			PORT:               s.port,
			ADDRESS:            s.address,
			HOSTED_BY_LOGO_URL: s.hostedByLogoURL,
			HOSTED_BY_TEXT:     s.hostedByText,
			TITLE_TEXT:         s.titleText,
		}

		err = view.ExecuteTemplate(w, file, e)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		http.FileServer(http.FS(index)).ServeHTTP(w, r)
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
