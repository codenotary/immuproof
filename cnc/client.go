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

package cnc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/spf13/viper"
	sdk "github.com/vchain-us/ledger-compliance-go/grpcclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCNCClient(lcApiKey, host, port, lcCertPath string, skipTlsVerify, noTls bool) (*sdk.LcClient, error) {
	if skipTlsVerify && noTls {
		return nil, errors.New("illegal parameters submitted: skip-tls-verify and no-tls arguments are both provided")
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.New("ledger compliance port is invalid")
	}

	currentOptions, err := getDialOptions(skipTlsVerify, noTls, lcCertPath)
	if err != nil {
		return nil, err
	}

	return sdk.NewLcClient(
		sdk.ApiKey(lcApiKey),
		sdk.MetadataPairs([]string{
			"version", "immuproof/0.1",
		}),
		sdk.Host(host),
		sdk.Port(p),
		sdk.Dir(viper.GetString("audit-state-folder")),
		sdk.DialOptions(currentOptions),
	), nil
}

func getDialOptions(skipTlsVerify, noTls bool, lcCertPath string) ([]grpc.DialOption, error) {
	currentOptions := []grpc.DialOption{}
	if !skipTlsVerify {
		if lcCertPath != "" {
			tlsCredentials, err := loadTLSCertificate(lcCertPath)
			if err != nil {
				return nil, fmt.Errorf("cannot load TLS credentials: %s", err)
			}
			currentOptions = append(currentOptions, grpc.WithTransportCredentials(tlsCredentials))
		} else {
			// automatic loading of local CA in os
			currentOptions = append(currentOptions, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
		}
	} else {
		currentOptions = append(currentOptions, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})))
	}
	if noTls {
		currentOptions = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}
	return currentOptions, nil
}

func loadTLSCertificate(certPath string) (credentials.TransportCredentials, error) {
	cert, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}
