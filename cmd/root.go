/*
Copyright © 2022 CodeNotary, Inc. All rights reserved

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/codenotary/immuproof/meta"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = NewRootCmd()

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "immuproof",
		Short: "Simple audit tool for CAS",
		Long: `Simple audit tool for CAS.

Environment variables:
  IMMUPROOF_API_KEY=
  IMMUPROOF_PORT=
  IMMUPROOF_HOST=
  IMMUPROOF_SKIP_TLS_VERIFY=
  IMMUPROOF_NO_TLS=
  IMMUPROOF_CERT=
  IMMUPROOF_HOST=
  IMMUPROOF_AUDIT_INTERVAL=
  IMMUPROOF_AUDIT_STATE_FOLDER=
  IMMUPROOF_STATE_HISTORY_SIZE=
  IMMUPROOF_STATE_HISTORY_FILE=
  IMMUPROOF_WEB_PORT=
  IMMUPROOF_WEB_ADDRESS=
  IMMUPROOF_WEB_KEY_FILE=
  IMMUPROOF_WEB_CERT_FILE=
  IMMUPROOF_WEB_HOSTED_BY_LOGO_URL=
  IMMUPROOF_WEB_HOSTED_BY_LOGO_LINK=
  IMMUPROOF_WEB_HOSTED_BY_TEXT=
  IMMUPROOF_WEB_TITLE_TEXT=


When immuproof is launched it fetches a fresh status from CAS backed by immudb and it verifies the integrity compared to an older one stored locally.
The idea is to check if previous state is "included" in the new one.
A rest service is also provided to allow the user to query the status of the audit.
A simple web UI is also provided to visualize data.

Eg:
# Collect 3 days of status checks (1 per hour) from CAS server
immuproof serve --api-key {your api-key} --port 443 --host admin.cas.codenotary.com --skip-tls-verify --audit-interval 1h --state-history-size 72
`,
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", meta.DefaultConfigFolder, fmt.Sprintf("config file (default is %s/.immuproof.yaml)", meta.DefaultConfigFolder))
	rootCmd.PersistentFlags().IntP("port", "p", meta.DefaultCNCPort, "CAS server port number")
	rootCmd.PersistentFlags().StringP("host", "a", meta.DefaultCNCHost, "CAS server host address")
	rootCmd.PersistentFlags().StringSlice("api-key", nil, "CAS api-keys. Can be specified multiple times. First key is used for signing. For each key provided related ledger is audit. If no key is provided, no audit is performed")
	rootCmd.PersistentFlags().String("cert", "", "local or absolute path to a certificate file needed to set up tls connection to a CAS server")
	rootCmd.PersistentFlags().Bool("skip-tls-verify", false, "disables tls certificate verification when connecting to a CAS server")
	rootCmd.PersistentFlags().Bool("no-tls", false, "allow insecure connections when connecting to a CAS server")

	viper.BindPFlags(rootCmd.PersistentFlags())

	viper.SetDefault("port", meta.DefaultCNCPort)
	viper.SetDefault("host", meta.DefaultCNCHost)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	err := ensureDir()
	cobra.CheckErr(err)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".immuproof" (without extension).
		viper.AddConfigPath(meta.DefaultConfigFolder)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".immuproof")
	}
	viper.SetEnvPrefix(strings.ToUpper(meta.AppName))
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func ensureDir() error {
	err := os.MkdirAll(meta.DefaultConfigFolder, 0755)
	if err != nil {
		return fmt.Errorf("failed to create config folder: %w", err)
	}
	err = os.MkdirAll(viper.GetString("audit-state-folder"), 0755)
	if err != nil {
		return fmt.Errorf("failed to create config folder: %w", err)
	}
	return nil
}
