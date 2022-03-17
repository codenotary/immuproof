package status

import (
	"sync"
	"time"
)

const Status_CORRUPTED_DATA = "CORRUPTED_DATA"
const Status_NORMAL = "NORMAL"

type StatusReport struct {
	SignerID      string    `json:"-"`
	Time          time.Time `json:"time"`
	Status        string    `json:"status"`
	PrevStateHash string    `json:"prev_state_hash"`
	NewStateHash  string    `json:"new_state_hash"`
}

type StatusReportMap struct {
	l sync.Mutex                `json:"-"`
	M map[string][]StatusReport `json:"status"`
}

func NewStatusReportMap() *StatusReportMap {
	return &StatusReportMap{
		M: make(map[string][]StatusReport),
	}
}

func (m *StatusReportMap) Add(report StatusReport) {
	m.l.Lock()
	defer m.l.Unlock()
	m.M[report.SignerID] = append(m.M[report.SignerID], report)
}

func (m *StatusReportMap) Get(signerID string) ([]StatusReport, bool) {
	m.l.Lock()
	defer m.l.Unlock()
	report, ok := m.M[signerID]
	return report, ok
}

func (m *StatusReportMap) Delete(signerID string) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.M, signerID)
}

func (m *StatusReportMap) GetAll() []StatusReport {
	m.l.Lock()
	defer m.l.Unlock()
	reports := make([]StatusReport, 0, len(m.M))
	for _, report := range m.M {
		reports = append(reports, report...)
	}
	return reports
}
