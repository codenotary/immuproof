package status

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStatusReportMap(t *testing.T) {
	historySize := 2
	m := NewStatusReportMap(historySize)
	sr1 := StatusReport{
		SignerID:      "s1",
		Time:          time.Time{},
		TimeZone:      "",
		Status:        "",
		PrevTxID:      0,
		PrevStateHash: "",
		NewTxID:       0,
		NewStateHash:  "",
	}

	sr2 := StatusReport{
		SignerID:      "s2",
		Time:          time.Time{},
		TimeZone:      "",
		Status:        "",
		PrevTxID:      0,
		PrevStateHash: "",
		NewTxID:       0,
		NewStateHash:  "",
	}

	m.Add(sr1)
	m.Add(sr2)

	msr := m.GetAllByLedger()

	require.Len(t, msr, 2)

	sr3 := StatusReport{
		SignerID:      "s2",
		Time:          time.Time{},
		TimeZone:      "",
		Status:        "",
		PrevTxID:      0,
		PrevStateHash: "",
		NewTxID:       0,
		NewStateHash:  "",
	}
	sr4 := StatusReport{
		SignerID:      "s2",
		Time:          time.Time{},
		TimeZone:      "",
		Status:        "",
		PrevTxID:      0,
		PrevStateHash: "",
		NewTxID:       0,
		NewStateHash:  "",
	}
	m.Add(sr3)
	m.Add(sr4)

	msr = m.GetAllByLedger()
	require.Len(t, msr, 2)
}
