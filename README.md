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

CAS staging environment

```shell
immuproof serve --api-key {your api key} --port 443 --host admin.cas-staging.codenotary.com --skip-tls-verify
```

## Others options

```yaml
Audit a ledger and launch an HTTP rest server to show audit results:

more detailed expl.

Usage:
  immuproof serve [flags]

Flags:
  -h, --help              help for serve
      --web-port string   rest server port (default "8091")

Global Flags:
      --api-key strings   Codenotary Cloud/CAS api-keys. Can be specified multiple times. First key is used for signing. For each key provided related ledger is audit. If no key is provided, no audit is performed
      --config string     config file (default is /home/falce/.config/ImmuProof/.immuproof.yaml)
  -a, --host string       Codenotary Cloud/CAS server host address (default "localhost")
      --lc-cert string    local or absolute path to a certificate file needed to set up tls connection to a Codenotary Cloud/CAS server
      --no-tls            allow insecure connections when connecting to a Codenotary Cloud/CAS server
  -p, --port int          Codenotary Cloud/CAS server port number (default 443)
      --skip-tls-verify   disables tls certificate verification when connecting to a Codenotary Cloud/CAS server
```
