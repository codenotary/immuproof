package rest

import (
	"encoding/json"
	"github.com/codenotary/immuproof/status"
	"log"
	"net/http"
)

type restServer struct {
	port      string
	statusMap *status.StatusReportMap
}

func NewRestServer(statusMap *status.StatusReportMap, port string) *restServer {
	return &restServer{statusMap: statusMap, port: port}
}

func (s *restServer) Serve() {
	http.HandleFunc("/", s.allStatuses)
	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}

func (s *restServer) allStatuses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.statusMap.GetAll())
}
