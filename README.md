# Rock Paper Scissors Lizard Spock

[![Go](https://img.shields.io/badge/Go-1.17-blue.svg)](https://golang.org)
[![build](https://github.com/arttet/rock-paper-scissors-lizard-spock/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/arttet/rock-paper-scissors-lizard-spock/actions/workflows/build.yml)
[![tests](https://github.com/arttet/rock-paper-scissors-lizard-spock/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/arttet/rock-paper-scissors-lizard-spock/actions/workflows/tests.yml)
[![codecov](https://codecov.io/gh/arttet/rock-paper-scissors-lizard-spock/branch/main/graph/badge.svg?token=7FHPNUZ1UZ)](https://codecov.io/gh/arttet/rock-paper-scissors-lizard-spock)
[![Swagger Validator](https://img.shields.io/swagger/valid/3.0?specUrl=https%3A%2F%2Fraw.githubusercontent.com%2Farttet%2Frock-paper-scissors-lizard-spock%2Fmain%2Fapi%2Fopenapi-spec%2Fapi.swagger.json)](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/arttet/rock-paper-scissors-lizard-spock/main/api/openapi-spec/api.swagger.json)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/arttet/rock-paper-scissors-lizard-spock/blob/main/LICENSE)

## Usage

1. Run Docker
```sh
docker-compose -f ./deploy/docker/docker-compose.yaml --project-directory ./ up -d
```

2. [Click me](https://arttet.github.io/rock-paper-scissors-lizard-spock)

## Services

### REST API

http://localhost:8080
- [/choices](http://localhost:8080/choices) - [GET] Get all the choices that are usable for the UI
- [/choice](http://localhost:8080/choice) - [GET] Get a randomly generated choice
- [/play](http://localhost:8080/play) - [POST] Play a round against a computer opponent

### gRPC

- http://localhost:8082

### Metrics

http://localhost:9100/metrics
- `http_microservice_requests_total`

### Status

http://localhost:8000
- [/live](http://localhost:8000/live) - Liveness Probe
- [/ready](http://localhost:8000/ready) - Readiness Probe
- [/version](http://localhost:8000/version) - Version and assembly information
- [/log/level](http://localhost:8000/log/level) - Show the current logging level
- [/swagger/{swagger_file}](http://localhost:8000/swagger/api.swagger.json) - Show a Swagger file
- [/debug/pprof](http://localhost:8000/debug/pprof) - Profiles

## UI

### Jaeger UI

Monitor and troubleshoot transactions in complex distributed systems.

- http://localhost:16686

### Prometheus UI

Prometheus is an open-source systems monitoring and alerting toolkit

- http://localhost:9090

### Swagger UI

The Swagger UI is an open source project to visually render documentation for an API defined with the OpenAPI (Swagger) Specification

- http://localhost:8008
