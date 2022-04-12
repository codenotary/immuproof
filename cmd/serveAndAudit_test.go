package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServeAndAudit(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"serve"})

	sa := NewServeAndAuditCmd()
	rootCmd.AddCommand(sa)
	err := rootCmd.Execute()
	require.Errorf(t, err, "no api-key provided")
}

func TestServeAndAuditFailConnectErr(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"serve", "--api-key", "signer-test", "--host", "wrong"})

	sa := NewServeAndAuditCmd()
	rootCmd.AddCommand(sa)
	err := rootCmd.Execute()
	require.Error(t, err)
}
