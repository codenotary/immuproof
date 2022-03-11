package cnc

import (
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/codenotary/immuproof/meta"
	sdk "github.com/vchain-us/ledger-compliance-go/grpcclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
	"strconv"
)

func NewCNCClient(lcApiKey, lcLedger, host, port, lcCertPath string, skipTlsVerify, noTls bool, signingPubKey *ecdsa.PublicKey) (*sdk.LcClient, error) {
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
			meta.LedgerHeaderName, lcLedger,
		}),
		sdk.Host(host),
		sdk.Port(p),
		sdk.Dir(meta.DefaultStateFolder),
		sdk.DialOptions(currentOptions),
		sdk.ServerSigningPubKey(signingPubKey),
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
