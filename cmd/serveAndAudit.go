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
		//todo: check args
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
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
	if err != nil {
		return err
	}
	err = client.Connect()
	if err != nil {
		return err
	}

	statusReportMap := status.NewStatusReportMap()
	simpleAuditor := audit.NewSimpleAuditor(client, statusReportMap)
	for _, a := range aks {
		simpleAuditor.AddApiKey(a)
	}
	restServer := rest.NewRestServer(statusReportMap, "8080")

	go restServer.Serve()
	go simpleAuditor.Audit()

	<-done
	return nil
}
