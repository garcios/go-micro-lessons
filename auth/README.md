# Authentication demo
This demo also uses the consul as the discovery service.

```go
    import (
  "github.com/go-micro/plugins/v4/registry/consul"
  ...
  )

    ...

        reg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("hello"),
		micro.WrapHandler(AuthMiddleware()),
	)
```

## Starting the consul server
```shell
podman run  -p 8500:8500 -p 8600:8600/udp --name=consul consul:v0.6.4 agent -server -bootstrap -ui -client=0.0.0.0
```

## Server
```shell
cd auth/server
go run main.go auth.go
```

## Client
```shell
cd auth
go run client/main.go
```