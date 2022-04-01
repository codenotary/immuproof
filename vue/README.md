# Immuproof UI


Install dependencies

```shell
npm install 
```
---


Build project
```shell
npm run build 
```
---

Build BE to compile files
```
go build -o immuproof main.go
```

---

Run with immuproof

```shell
./immuproof serve --api-key {apiKey} --port 443  --host admin.cas-staging.codenotary.com --skip-tls-verify --audit-interval 1h --state-history-size 720 
```


---
##For working locally with Vue;

First run backend for API connection

Without BE running, FE will not display datas!

```shell
./immuproof serve --api-key {apiKey} --port 443  --host admin.cas-staging.codenotary.com --skip-tls-verify --audit-interval 1h --state-history-size 720 
```

```shell
npm run serve
```