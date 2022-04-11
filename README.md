# Immuproof

[![Coverage Status](https://coveralls.io/repos/github/codenotary/immuproof/badge.svg)](https://coveralls.io/github/codenotary/immuproof)

CAS Validation Service.

When immuproof is launched it fetches a fresh status from [immudb](https://github.com/codenotary/immudb) the immutable database [CAS](https://cas.codenotary.com) is build on and it verifies the integrity compared to an older one stored locally.
The validation service checks if the previous state is "included" in the new state of immudb.
A REST service is also provided to allow the user to query the status of the validation as well as a Web UI to visualize data.

## Golang version

Currently supported Go version is `1.17`

## Build

```shell
go build -o immuproof main.go
```

## Usage

Local environment

```shell
immuproof serve --api-key {your CAS api key} --port 3324 --no-tls
```

CAS environment

```shell
immuproof serve --api-key {your CAS api key} --port 443 --host cas.codenotary.com 
```

## Usage with docker

```shell
docker pull codenotary/immuproof:latest
docker run -p 8091:8091 codenotary/immuproof serve --api-key {your api key} --port 443 --host cas.codenotary.com --audit-interval 1h --state-history-size 72
```

In order to keep the audit history and [immudb](https://github.com/codenotary/immudb) status file it's recommended to run the service with a mounted volume inside the docker container using following flags:

```shell
--audit-state-folder={mountpoint inside container}
--state-history-file={mountpoint inside container/filename}
```

or environment variables:

```shell
IMMUPROOF_AUDIT_STATE_FOLDER={mountpoint inside container}
IMMUPROOF_STATE_HISTORY_FILE={mountpoint inside container/filename}
```

## HTTPS

Following commands can be used to generate a self-signed certificate for the local server.

```shell
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

Launch immuproof with the generated certificate:

```shell
immuproof serve --api-key {your CAS api key} --port 443 --host cas.codenotary.com --audit-interval 1s --state-history-size 72 --web-cert-file server.crt --web-key-file server.key
```

## Environment variables
```shell
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
  IMMUPROOF_WEB_TITLE_TEXT=
  IMMUPROOF_WEB_HOSTED_BY_LOGO_URL=
  IMMUPROOF_WEB_HOSTED_BY_TEXT=
```

## Others serve options

```yaml
Audit a ledger and launch an HTTP rest server to show audit results.

Eg:
# Collect 3 days of status checks (1 per hour) from CAS server
  immuproof serve --api-key {your CAS api-key} --port 443 --host cas.codenotary.com --audit-interval 1h --state-history-size 72

Usage:
  immuproof serve [flags]

Flags:
  --audit-interval duration         interval between audit runs (default 1h0m0s)
  --audit-state-folder string       folder to store immudb immutable state (default "HOME/.local/state/immuproof")
  -h, --help                            help for serve
  --state-history-file string       absolute file path to store history of immutable states. (JSON format) (default "HOME/.local/state/immuproof/state-history.json")
  --state-history-size int          max size of the history of immutable states. (default 90)
  --web-address string              rest server address (default "localhost")
  --web-cert-file string            certificate file absolute path
  --web-hosted-by-logo-url string   URL to hosted by logo
  --web-hosted-by-text string       displayed subtitle for hosted by logo (default "Hosted by:")
  --web-key-file string             key file absolute path
  --web-port string                 rest server port (default "8091")
  --web-title-text string           displayed title text (default "COMMUNITY ATTESTATION SERVICE VALIDATOR")

Global Flags:
  --api-key strings   CAS api-keys. Can be specified multiple times. First key is used for signing. For each key provided related ledger is audit. If no key is provided, no audit is performed
  --cert string       local or absolute path to a certificate file needed to set up tls connection to a CAS server
  --config string     config file (default is /root/.config/immuproof/.immuproof.yaml) (default "HOME/.config/immuproof")
  -a, --host string       CAS server host address (default "localhost")
  --no-tls            allow insecure connections when connecting to a CAS server
  -p, --port int          CAS server port number (default 443)
  --skip-tls-verify   disables tls certificate verification when connecting to a CAS server
```
