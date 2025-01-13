# Go micro lessons
This comprises key insights and crucial programs that have immensely contributed to my understanding of the Go-Micro 
framework.

## Install binaries

Go-Micro V4
```shell
go get go-micro.dev/v4@latest
```


Protobuf Compiler
```shell
brew install protobuf
```

Protobuf Go Plugin
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
```

Micro CLI
```shell
go install github.com/asim/go-micro/cmd/micro@latest
```

Open API
```shell
go get github.com/google/gnostic
go install github.com/google/gnostic/cmd/protoc-gen-openapi
```

Dashboard
```shell
go install github.com/go-micro/dashboard@latest
```


## mod file
```go
module github.com/oskiegarcia/go-micro-lessons

go 1.22.2

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/go-micro/plugins/v4/broker/kafka v1.2.0
    github.com/go-micro/plugins/v4/broker/mqtt v1.2.0
    github.com/go-micro/plugins/v4/client/grpc v1.2.1
    github.com/go-micro/plugins/v4/registry/etcd v1.2.0
    github.com/go-micro/plugins/v4/server/http v1.2.2
    github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry v1.2.0
    github.com/go-micro/plugins/v4/wrapper/trace/opentracing v1.2.0
    github.com/google/uuid v1.6.0
    github.com/opentracing/opentracing-go v1.2.0
    github.com/uber/jaeger-client-go v2.30.0+incompatible
    github.com/xpunch/go-micro-example/v4 v4.0.0-20230924144533-5ab2063c35a3
    go-micro.dev/v4 v4.11.0
    go.opentelemetry.io/otel v1.33.0
    go.opentelemetry.io/otel/exporters/jaeger v1.17.0
    go.opentelemetry.io/otel/sdk v1.33.0
    google.golang.org/protobuf v1.36.2
)
```



### References:
- https://micro.dev/getting-started
- https://github.com/micro/services
- https://github.com/xpunch/go-micro-example/blob/main/v4/README.md
- https://www.twilio.com/en-us/blog/a-practical-guide-to-creating-microservices-with-go-micro