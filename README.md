# Immuproof

Simple audit tool for CAS and CodeNotaryCloud services.

When immuproof is launched it fetches a fresh status from CodeNotaryCloud or CAS backed by immudb and it verifies the integrity compared to an older one stored locally.
The idea is to check if previous state is "included" in the new one.
A REST service is also provided to allow the user to query the status of the audit.
A simple web UI is also provided to visualize data.

## Build

```shell
go build -o immuproof main.go
```

## Usage

Local environment

```shell
immuproof serve --api-key {your api key} --port 3324 --no-tls
```

CAS environment

```shell
immuproof serve --api-key {your api key} --port 443 --host admin.cas.codenotary.com --skip-tls-verify
```

## HTTPS

Following commands can be used to generate a self-signed certificate for the local server.

```shell
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

Launch immuproof with the generated certificate:

```shell
immuproof serve --api-key {your api key} --port 443 --host admin.cas.codenotary.com --audit-interval 1s --state-history-size 72 --web-cert-file server.crt --web-key-file server.key

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
```

## Others serve options

```yaml
Audit a ledger and launch an HTTP rest server to show audit results.

Eg:
# Collect 3 days of status checks (1 per hour) from CAS server
  immuproof serve --api-key {your api-key} --port 443 --host admin.cas.codenotary.com --skip-tls-verify --audit-interval 1h --state-history-size 72

Usage:
  immuproof serve [flags]

Flags:
  --audit-interval duration     interval between audit runs (default 1h0m0s)
  --audit-state-folder string   folder to store immudb immutable state (default "/home/falce/.local/state/immuproof")
  -h, --help                        help for serve
  --state-history-file string   absolute file path to store history of immutable states. (JSON format) (default "/home/falce/.local/state/immuproof/state-history.json")
  --state-history-size int      max size of the history of immutable states. (default 90)
  --web-address string          rest server address (default "localhost")
  --web-cert-file string        certificate file absolute path
  --web-key-file string         key file absolute path
  --web-port string             rest server port (default "8091")

Global Flags:
  --api-key strings   Codenotary Cloud/CAS api-keys. Can be specified multiple times. First key is used for signing. For each key provided related ledger is audit. If no key is provided, no audit is performed
  --cert string       local or absolute path to a certificate file needed to set up tls connection to a Codenotary Cloud/CAS server
  --config string     config file (default is /home/falce/.config/immuproof/.immuproof.yaml) (default "/home/falce/.config/immuproof")
  -a, --host string       Codenotary Cloud/CAS server host address (default "localhost")
  --no-tls            allow insecure connections when connecting to a Codenotary Cloud/CAS server
  -p, --port int          Codenotary Cloud/CAS server port number (default 443)
  --skip-tls-verify   disables tls certificate verification when connecting to a Codenotary Cloud/CAS server
```
