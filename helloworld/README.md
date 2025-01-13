# Hello World
This example demonstrates how to write grpc server and client.
> 1. Generate protobuf files.
> 2. Implement your service, and register with micro service.

## Protobuf
### Change directory
```shell
cd <go-micro-lessons>
```

### Generate protobuf files
```
protoc --go_out=proto --micro_out=proto proto/helloworld.proto
```


## GRPC

### Change directory
```shell
cd <go-micro-lessons>/helloworld
```


### Run Server
```
go run server/main.go --server grpc --server_name helloworld.srv
```

### Run Client
```
go run client/main.go --client grpc
```