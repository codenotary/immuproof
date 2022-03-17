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
	"github.com/codenotary/immuproof/audit"
	"github.com/codenotary/immuproof/cnc"
	"github.com/codenotary/immuproof/meta"
	"github.com/codenotary/immuproof/rest"
	"github.com/codenotary/immuproof/status"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Audit a ledger and launch an HTTP rest server to show audit results",
	Long: `Audit a ledger and launch an HTTP rest server to show audit results:

more detailed expl.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		return ServeAndAudit()
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if viper.GetBool("lc-skip-tls-verify") && viper.GetBool("no-tls") {
			return fmt.Errorf("--lc-skip-tls-verify and --no-tls are mutually exclusive")
		}
		return nil
	},
}

func init() {
	serveCmd.Flags().String("web-port", "8091", "rest server port")
	serveCmd.Flags().Duration("audit-interval", meta.DefaultAuditInterval, "interval between audit runs")
	serveCmd.Flags().String("audit-state-folder", meta.DefaultStateFolder, "folder to store immudb immutable state")
	serveCmd.Flags().String("state-cache-file", meta.DefaultStateMapFileName, "absolute file path to store history of immutable states. (JSON format)")
	rootCmd.AddCommand(serveCmd)
	viper.BindPFlags(serveCmd.Flags())
}

func ServeAndAudit() error {

	done := make(chan bool)

	aks := viper.GetStringSlice("api-key")
	if len(aks) == 0 {
		return fmt.Errorf("no api-key provided")
	}
	client, err := cnc.NewCNCClient(
		aks[0],
		viper.GetString("host"),
		viper.GetString("port"),
		viper.GetString("lc-cert"),
		viper.GetBool("lc-skip-tls-verify"),
		viper.GetBool("no-tls"),
	)
	cobra.CheckErr(err)
	cobra.CheckErr(client.Connect())

	statusReportMap := status.NewStatusReportMap()
	simpleAuditor := audit.NewSimpleAuditor(client, statusReportMap, viper.GetDuration("audit-interval"))
	for _, a := range aks {
		simpleAuditor.AddApiKey(a)
	}
	restServer := rest.NewRestServer(statusReportMap, viper.GetString("web-port"))

	go func() {
		cobra.CheckErr(restServer.Serve())
	}()
	go func() {
		cobra.CheckErr(simpleAuditor.Audit())
	}()

	<-done
	return nil
}
