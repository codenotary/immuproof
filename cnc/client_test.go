package cnc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCNCClient(t *testing.T) {
	lcApiKey := "test_api_key"
	host := "test_host"
	port := "8080"
	lcCertPath := ""
	var skipTlsVerify, noTls bool
	c, err := NewCNCClient(lcApiKey, host, port, lcCertPath, skipTlsVerify, noTls)
	require.NoError(t, err)
	if c == nil {
		t.Error("NewCNCClient() returned nil")
	}
}

func TestNewCNCClientTLS(t *testing.T) {
	lcApiKey := "test_api_key"
	host := "test_host"
	port := "8080"
	lcCertPath := "./cnctest/test_server.crt"
	var skipTlsVerify, noTls bool
	c, err := NewCNCClient(lcApiKey, host, port, lcCertPath, skipTlsVerify, noTls)
	require.NoError(t, err)
	if c == nil {
		t.Error("NewCNCClient() returned nil")
	}

	lcCertPath = "./cnctest/wrong.crt"
	_, err = NewCNCClient(lcApiKey, host, port, lcCertPath, skipTlsVerify, noTls)
	require.Error(t, err)

	lcCertPath = "./cnctest/missing.crt"
	_, err = NewCNCClient(lcApiKey, host, port, lcCertPath, skipTlsVerify, noTls)
	require.Error(t, err)
}

func TestNewCNCClientNOTLS(t *testing.T) {
	lcApiKey := "test_api_key"
	host := "test_host"
	port := "8080"
	lcCertPath := "./cnctest/test_server.crt"

	c, err := NewCNCClient(lcApiKey, host, port, lcCertPath, true, true)
	require.Contains(t, err.Error(), "illegal parameters submitted: skip-tls-verify and no-tls arguments are both provided")

	c, err = NewCNCClient(lcApiKey, host, "wrong", lcCertPath, false, false)
	require.Contains(t, err.Error(), "ledger compliance port is invalid")

	c, err = NewCNCClient(lcApiKey, host, port, lcCertPath, false, true)
	require.NoError(t, err)
	if c == nil {
		t.Error("NewCNCClient() returned nil")
	}
}
