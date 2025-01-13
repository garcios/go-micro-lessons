# Authentication demo
This demo also uses the consul as the discovery service.


## Building a Custom Service Discovery 
One of the key features of Go-Micro is its pluggable service discovery mechanism. By default, Go-Micro uses the 
micro/go-micro/registry/mdns package for service discovery, which is based on multicast DNS. However, you can easily 
replace it with your own implementation or use a third-party package.

To demonstrate this, let’s replace the default mDNS-based service discovery with Consul, a popular service mesh 
solution. First, install the github.com/go-micro/plugins/v4/registry/consul package:

```shell
go get -u github.com/go-micro/plugins/v4/registry/consul
```


Next, update your `main` file to use the Consul registry:
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

Now, your service will register itself with a local Consul agent. Make sure you have Consul installed and running on 
your machine before starting the service.

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