package rest

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/codenotary/immuproof/status"
	"github.com/stretchr/testify/require"
)

func TestRest(t *testing.T) {
	srm := status.NewStatusReportMap(2)
	rs, err := NewRestServer(srm, "0", "", "", "", "", "", "", "")
	require.NoError(t, err)
	require.IsType(t, &restServer{}, rs)
}

func TestStatusHandler_ServeHTTP(t *testing.T) {
	srm := status.NewStatusReportMap(2)
	rs, err := NewRestServer(srm, "0", "", "", "", "", "", "", "")
	require.NoError(t, err)
	require.IsType(t, &restServer{}, rs)
	req := httptest.NewRequest("GET", "/", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	sh := statusHandler{statusMap: srm}
	sh.ServeHTTP(rr, req)
	resp := rr.Result()
	require.Equal(t, 200, resp.StatusCode)
}

func TestCountHandler_ServeHTTP(t *testing.T) {
	srm := status.NewStatusReportMap(2)
	sr1 := status.StatusReport{
		SignerID: "s2",
	}
	sr2 := status.StatusReport{
		SignerID: "s2",
	}

	srm.Add(sr1)
	srm.Add(sr2)

	rs, err := NewRestServer(srm, "0", "", "", "", "", "", "", "")
	require.NoError(t, err)
	require.IsType(t, &restServer{}, rs)
	req := httptest.NewRequest("GET", "/", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	sh := countHandler{statusMap: srm}
	sh.ServeHTTP(rr, req)
	resp := rr.Result()
	require.Equal(t, 200, resp.StatusCode)
}

func TestWebHandler_ServeHTTP(t *testing.T) {
	srm := status.NewStatusReportMap(2)
	rs, err := NewRestServer(srm, "0", "", "", "", "", "", "", "")
	require.NoError(t, err)
	require.IsType(t, &restServer{}, rs)
	req := httptest.NewRequest("GET", "/", nil)
	req.URL.Path = "/js/app.js"
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	sh := webHandler{}
	sh.ServeHTTP(rr, req)
	resp := rr.Result()
	require.Equal(t, 200, resp.StatusCode)
	req2 := httptest.NewRequest("GET", "/", nil)
	sh.ServeHTTP(rr, req2)
	resp2 := rr.Result()
	require.Equal(t, 200, resp2.StatusCode)
}

func TestServer_ServeAndStop(t *testing.T) {
	srm := status.NewStatusReportMap(2)
	rs, err := NewRestServer(srm, "0", "", "", "", "", "", "", "")
	require.NoError(t, err)
	require.IsType(t, &restServer{}, rs)
	go func() {
		rs.Serve()
	}()
	time.Sleep(time.Second)
	err = rs.Stop()
	require.NoError(t, err)
}
