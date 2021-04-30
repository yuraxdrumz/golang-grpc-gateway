# Grpc Gateway

An example on how to add an http gateway infront of a grpc server

## Requirements

- Latest protoc in $PATH

## Usage

1. cURL <http://localhost:8081/api/v1/{name}>

## Transcoding HTTP/JSON to gRPC

<https://cloud.google.com/endpoints/docs/grpc/transcoding>

## Protobuf Files Generation (Grpc service + messages + gateway)

```bash
protoc -I ./proto --go_out ./generated --go_opt paths=source_relative --go-grpc_out ./generated --go-grpc_opt paths=source_relative --grpc-gateway_out ./generated --grpc-gateway_opt paths=source_relative ./proto/*.proto
```
