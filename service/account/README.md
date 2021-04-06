# Account Service

This is the Account service

Generated with

```
micro new github.com/kmaguswira/micro-clean/service/account --namespace=kmaguswira --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: account
- Type: srv
- Alias: account

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

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
./account-srv
```

Build a docker image
```
make docker
```