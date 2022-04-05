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
	lcCertPath := "./cnctest/server.crt"
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
