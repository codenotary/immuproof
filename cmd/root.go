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
	"github.com/codenotary/immuproof/meta"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "immuproof",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is %s/.immuproof.yaml)", meta.DefaultConfigFolder))
	rootCmd.PersistentFlags().IntP("port", "p", meta.DefaultCNCPort, "Codenotary Cloud server port number")
	rootCmd.PersistentFlags().StringP("address", "a", meta.DefaultCNCHost, "Codenotary Cloud server host address")
	rootCmd.PersistentFlags().String("api-key", "", "cnc api-key")
	rootCmd.PersistentFlags().String("lc-cert", "", "local or absolute path to a certificate file needed to set up tls connection to a Codenotary Cloud server")
	rootCmd.PersistentFlags().Bool("skip-tls-verify", false, "disables tls certificate verification when connecting to a Codenotary Cloud server")
	rootCmd.PersistentFlags().Bool("no-tls", false, "allow insecure connections when connecting to a Codenotary Cloud server")

	viper.BindPFlags(rootCmd.PersistentFlags())

	viper.SetDefault("port", meta.DefaultCNCPort)
	viper.SetDefault("address", meta.DefaultCNCHost)
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
		return err
	}
	err = os.MkdirAll(meta.DefaultStateFolder, 0755)
	if err != nil {
		return err
	}
	return nil
}
