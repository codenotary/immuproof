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

	"github.com/rs/cors"

	"github.com/codenotary/immuproof/status"
)

//go:embed internal/embed
var content embed.FS

type restServer struct {
	port          string
	statusHandler *statusHandler
}

type statusHandler struct {
	port      string
	statusMap *status.StatusReportMap
}

func NewRestServer(statusMap *status.StatusReportMap, port string) *restServer {
	return &restServer{
		port: port,
		statusHandler: &statusHandler{
			port:      port,
			statusMap: statusMap,
		},
	}
}

func (s *restServer) Serve() error {
	log.Printf("Starting REST server on port %s", s.port)
	log.Print("UI is exposed on /")
	log.Print("REST server is exposed on /api/status")

	mux := http.NewServeMux()

	muxCors := cors.Default().Handler(mux)

	index, err := fs.Sub(content, "internal/embed")
	if err != nil {
		return err
	}

	mux.Handle("/", http.FileServer(http.FS(index)))
	mux.Handle("/api/status", s.statusHandler)

	return http.ListenAndServe(fmt.Sprintf(":%s", s.port), muxCors)
}

func (s *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.statusMap.GetAll())
}
