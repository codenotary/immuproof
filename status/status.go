package status

import (
	"sync"
	"time"
)

const Status_CORRUPTED_DATA = "CORRUPTED_DATA"
const Status_NORMAL = "NORMAL"

type StatusReport struct {
	SignerID string
	Time     time.Time
	Status   string
}

type StatusReportMap struct {
	l sync.Mutex
	m map[string]StatusReport
}

func NewStatusReportMap() *StatusReportMap {
	return &StatusReportMap{
		m: make(map[string]StatusReport),
	}
}

func (m *StatusReportMap) Add(report StatusReport) {
	m.l.Lock()
	defer m.l.Unlock()
	m.m[report.SignerID] = report
}

func (m *StatusReportMap) Get(signerID string) (StatusReport, bool) {
	m.l.Lock()
	defer m.l.Unlock()
	report, ok := m.m[signerID]
	return report, ok
}

func (m *StatusReportMap) Delete(signerID string) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.m, signerID)
}

func (m *StatusReportMap) GetAll() []StatusReport {
	m.l.Lock()
	defer m.l.Unlock()
	reports := make([]StatusReport, 0, len(m.m))
	for _, report := range m.m {
		reports = append(reports, report)
	}
	return reports
}
