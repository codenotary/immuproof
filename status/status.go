package status

import (
	"container/heap"
	"sync"
	"time"

	"github.com/spf13/viper"
)

const Status_CORRUPTED_DATA = "CORRUPTED_DATA"
const Status_NORMAL = "NORMAL"
const Status_UNKNOWN = "UNKNOWN"

type StatusReport struct {
	SignerID      string    `json:"signer_id"`
	Time          time.Time `json:"time"`
	Status        string    `json:"status"`
	PrevTxID      uint64    `json:"prev_tx_id"`
	PrevStateHash string    `json:"prev_state_hash"`
	NewTxID       uint64    `json:"new_tx_id"`
	NewStateHash  string    `json:"new_state_hash"`
}

type StatusReportMap struct {
	l sync.Mutex                 `json:"-"`
	M map[string]*statusReportPQ `json:"status"`
}

func NewStatusReportMap() *StatusReportMap {
	return &StatusReportMap{
		M: make(map[string]*statusReportPQ),
	}
}

func (m *StatusReportMap) Add(report StatusReport) {
	m.l.Lock()
	defer m.l.Unlock()

	var pq *statusReportPQ
	var ok bool
	if pq, ok = m.M[report.SignerID]; !ok {
		ppq := make(statusReportPQ, 5)
		pq = &ppq
	}
	for pq.Len() > viper.GetInt("state-cache-size") {
		heap.Pop(pq)
	}
	heap.Push(pq, &report)
}

func (m *StatusReportMap) GetAll() []*StatusReport {
	m.l.Lock()
	defer m.l.Unlock()
	reports := make([]*StatusReport, 0, len(m.M))
	for _, report := range m.M {
		reports = append(reports, report.GetAll()...)
	}
	return reports
}

type statusReportPQ []*StatusReport

func (pq statusReportPQ) Len() int           { return len(pq) }
func (pq statusReportPQ) Less(i, j int) bool { return pq[i].Time.Before(pq[j].Time) }
func (pq statusReportPQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *statusReportPQ) Push(x interface{}) {
	*pq = append(*pq, x.(*StatusReport))
}

func (pq *statusReportPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (pq *statusReportPQ) GetAll() []*StatusReport {
	return *pq
}
