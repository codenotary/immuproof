package rest

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/codenotary/immuproof/status"
	"io/fs"
	"log"
	"net/http"
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

	mutex := http.NewServeMux()
	index, err := fs.Sub(content, "internal/embed")
	if err != nil {
		return err
	}
	mutex.Handle("/", http.FileServer(http.FS(index)))
	mutex.Handle("/api/status", s.statusHandler)
	return http.ListenAndServe(fmt.Sprintf(":%s", s.port), mutex)
}

func (s *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.statusMap.GetAll())
}
