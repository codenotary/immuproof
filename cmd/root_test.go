package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoot(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})
	err := rootCmd.Execute()
	require.NoError(t, err)
}
