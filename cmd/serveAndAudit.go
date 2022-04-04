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
	Long: `Audit a ledger and launch an HTTP rest server to show audit results.

Eg:
# Collect 3 days of status checks (1 per hour) from CAS server
immuproof serve --api-key {your api-key} --port 443 --host admin.cas.codenotary.com --skip-tls-verify --audit-interval 1h --state-history-size 72
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return ServeAndAudit()
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if viper.GetBool("skip-tls-verify") && viper.GetBool("no-tls") {
			return fmt.Errorf("--skip-tls-verify and --no-tls are mutually exclusive")
		}
		if (viper.IsSet("web-cert-file") || viper.IsSet("web-key-file")) && (!(viper.IsSet("web-cert-file") && viper.IsSet("web-key-file"))) {
			return fmt.Errorf("--web-cert-file and --web-key-file must be used together")
		}
		return nil
	},
}

func init() {
	serveCmd.Flags().String("web-port", "8091", "rest server port")
	serveCmd.Flags().String("web-address", "localhost", "rest server address")
	serveCmd.Flags().String("web-cert-file", "", "certificate file absolute path")
	serveCmd.Flags().String("web-key-file", "", "key file absolute path")
	serveCmd.Flags().String("web-hosted-by-logo-url", "", "URL to hosted by logo")
	serveCmd.Flags().String("web-hosted-by-text", "Hosted by:", "displayed subtitle for hosted by logo")
	serveCmd.Flags().String("web-title-text", "COMMUNITY ATTESTATION SERVICE STATUS PAGE", "displayed title text")
	serveCmd.Flags().Duration("audit-interval", meta.DefaultAuditInterval, "interval between audit runs")
	serveCmd.Flags().String("audit-state-folder", meta.DefaultStateFolder, "folder to store immudb immutable state")
	serveCmd.Flags().Int("state-history-size", 90, "max size of the history of immutable states.")
	serveCmd.Flags().String("state-history-file", meta.DefaultStateMapFileName, "absolute file path to store history of immutable states. (JSON format)")
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
		viper.GetString("cert"),
		viper.GetBool("skip-tls-verify"),
		viper.GetBool("no-tls"),
	)
	cobra.CheckErr(err)
	cobra.CheckErr(client.Connect())

	statusReportMap := status.NewStatusReportMap()
	simpleAuditor := audit.NewSimpleAuditor(client, statusReportMap, viper.GetDuration("audit-interval"))
	for _, a := range aks {
		simpleAuditor.AddApiKey(a)
	}
	restServer := rest.NewRestServer(statusReportMap,
		viper.GetString("web-port"),
		viper.GetString("web-address"),
		viper.GetString("web-cert-file"),
		viper.GetString("web-key-file"),
		viper.GetString("web-hosted-by-logo-url"),
		viper.GetString("web-hosted-by-text"),
		viper.GetString("web-title-text"))

	go func() {
		cobra.CheckErr(restServer.Serve())
	}()
	go func() {
		cobra.CheckErr(simpleAuditor.Start())
	}()

	<-done

	simpleAuditor.Stop()

	return nil
}
