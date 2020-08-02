# Helloworld Service

This is the Helloworld service

Generated with

```
micro new --namespace=kzmake.example --alias=helloworld --type=service --plugin=registry=etcd:broker=rabbitmq github.com/kzmake/micro-helloworld
```

## Getting Started

- [Helloworld Service](#helloworld-service)
  - [Getting Started](#getting-started)
  - [Configuration](#configuration)
  - [Dependencies](#dependencies)
  - [Usage](#usage)

## Configuration

- FQDN: kzmake.example.service.helloworld
- Type: service
- Alias: helloworld

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./helloworld-service --registry etcd
```

Build a docker image
```
make docker
```
