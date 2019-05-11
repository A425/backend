# Dataservice Service

This is the Dataservice service

Generated with

```
micro new backend/dataservice --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.dataservice
- Type: srv
- Alias: dataservice

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./dataservice-srv
```

Build a docker image
```
make docker
```