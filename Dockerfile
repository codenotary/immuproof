FROM golang:1.17 as build
WORKDIR /src
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o immuproof main.go

FROM debian:bullseye-slim as bullseye-slim
LABEL org.opencontainers.image.authors="CodeNotary, Inc. <info@codenotary.com>"
RUN apt-get update && apt-get install ca-certificates -y
COPY --from=build /src/immuproof /usr/sbin/immuproof

EXPOSE 8091

ENTRYPOINT ["/usr/sbin/immuproof", "serve"]
