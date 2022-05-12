<!---
This file is auto-generate by a github hook please modify README.md.tpl if you don't want to loose your work
-->
# birdwatch
[![GitHub issues](https://img.shields.io/github/issues/Stalis/birdwatch?style=flat-square)](https://github.com/Stalis/birdwatch/issues)
[![GitHub license](https://img.shields.io/github/license/Stalis/birdwatch?style=flat-square)](https://github.com/Stalis/birdwatch/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/Stalis/birdwatch)](https://goreportcard.com/report/github.com/Stalis/birdwatch)
![Gitlab pipeline status](https://img.shields.io/gitlab/pipeline-status/Stalis/birdwatch?branch=6-fill-readmemd&label=buid%206-fill-readmemd&style=flat-square)

Monitoring service with gRPC api

## Requirements
- OS:
  - macOS (tested on 12 `amd64` and `arm64`)
  - Linux (tested on Fedora 35 kernel version `5.17.4`)

## Install

### From binaries
Download it from github releases

### Docker image
```bash
git clone https://github.com/Stalis/birdwatch
cd birdwatch
make docker
```
building docker image with tag `birdwatch`

### Source build
Run this commands
```bash
git clone https://github.com/Stalis/birdwatch
cd birdwatch
make build
```
Builded server in `bin/server`

## Configure
Service can be configurated by environment variables, cli args and config file
```yaml
# config.yaml
host: 0.0.0.0 # Host name for listening (env: BW_HOST, arg: -h,--host)
port: 50051   # Api server port (env: BW_PORT, arg: -p, --port)
logging:      # Logging settings
  verbose: true      # Print human-readable logs to stdout (env: BW_LOGGING_VERBOSE, arg: -v, --logging.verbose)
  level: Info        # Logging level, one of [Debug, Info, Warn, Error] (env)
  file: ./server.log # File for logging with json-format
memory:       # Memory stats scanner settings
  enable: true       # If false - disabling scanner
  interval: 1s       # Interval for data scanning
```

Name of argc and env vars are similiar to config var names:
environment vars are `UPPER_CASE` with prefix `BW_`, args are yaml-path for param, e.g. `--memory.enable`

### Full list of params

| config.yaml     | ENV                | CLI arg                   |
|-----------------|--------------------|---------------------------|
| host            | BW_HOST            | `-h`, `--host`            |
| port            | BW_PORT            | `-p`, `--port`            |
| logging.verbose | BW_LOGGING_VERBOSE | `-v`, `--logging.verbose` |
| logging.level   | BW_LOGGING_LEVEL   | `--logging.level`         |
| logging.file    | BW_LOGGING_FILE    | `--logging.level`         |
| memory.enable   | BW_MEMORY_ENABLE   | `--memory.enable`         |
| memory.interval | BW_MEMORY_INTERVAL | `--memory.interval`       |

## Development

For changing API, install `protoc`([install instructions](https://grpc.io/docs/protoc-installation/)) and `protoc-gen-go` ([install instructions](https://grpc.io/docs/languages/go/quickstart/))

For linting install `golangci-lint`([install instructions](https://golangci-lint.run/usage/install/#local-installation))

Before creating PR run `make lint` and `make test`, pipeline will not aprove it without passing linting and testing