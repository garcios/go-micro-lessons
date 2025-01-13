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
    github.com/patrickmn/go-cache v2.1.0+incompatible
    github.com/uber/jaeger-client-go v2.30.0+incompatible
    github.com/xpunch/go-micro-example/v4 v4.0.0-20230924144533-5ab2063c35a3
    go-micro.dev/v4 v4.11.0
    go.opentelemetry.io/otel v1.33.0
    go.opentelemetry.io/otel/exporters/jaeger v1.17.0
    go.opentelemetry.io/otel/sdk v1.33.0
    google.golang.org/protobuf v1.36.2
)
```

## Go-Micro vs Micro
### Go-Micro
An open-source framework for building microservices in Go. It provides tools and modules to abstract distributed systems
challenges, and includes features like automated service discovery and client-side load balancing. Go-micro uses a 
pluggable architecture, which means that default components can be replaced with other components that conform to the 
underlying interface definitions.

### Micro
A microservice ecosystem that includes go-micro. In version 3, micro simplified the development model and consolidated 
the micro and go-micro libraries. This means that developers only need to use the libraries offered in micro. Micro 
comes with a built-in Go framework for service-based development.

### Comparison
In this Markdown document, we will provide a comparison between Go Micro and Micro, highlighting the key differences 
between the two. Go Micro and Micro are both frameworks for building distributed systems and microservices.

1. __Extensibility__: Go Micro provides a highly extensible framework that enables developers to easily add their own 
functionality and customize the system according to their requirements. On the other hand, Micro offers a more 
opinionated approach, providing pre-defined components and plugins to assist in building microservices, which may 
limit the level of extensibility.

2. __Language Support__: Go Micro is primarily focused on supporting applications written in the Go programming 
language. It provides a Go-centric approach to building microservices, leveraging Go's features and ecosystem. 
In contrast, Micro aims to be language-agnostic, allowing developers to build microservices using different programming 
languages such as Go, Python, and Node.js.

3. __Service Discovery__: Go Micro utilizes service discovery to handle the registration and discovery of microservices 
within a distributed system. It includes support for various service discovery mechanisms like DNS, Consul, and etcd. 
Micro, on the other hand, relies on its own built-in service discovery system called "mDNS" (multicast DNS), which 
enables decentralized service discovery without the need for external dependencies.

4. __Transport Protocols__: Go Micro provides support for multiple transport protocols, including HTTP, TCP, and NATS. 
It allows developers to choose the most suitable transport protocol based on their use case. In contrast, Micro 
emphasizes the use of the HTTP protocol as the primary transport mechanism for communication between microservices.

5. __API Gateway__: Go Micro does not include an API gateway as part of its core functionality. However, it provides the
flexibility to build and integrate an API gateway within the system as per the developer's requirements. On the other 
hand, Micro offers a built-in API gateway called "Micro API" that allows developers to easily expose microservices as 
HTTP APIs without additional configuration.

6. __Error Handling__: Go Micro focuses on providing a pluggable error handling mechanism, allowing developers to define
their own error handling strategies for different scenarios. Micro, on the other hand, provides a standardized error 
handling mechanism that follows the microservices best practices, promoting a consistent error handling approach across 
services.

In Summary, Go Micro offers a more extensible and Go-centric approach with support for multiple transport protocols and
flexible error handling, while Micro provides a more opinionated and language-agnostic framework with its own built-in 
service discovery system and API gateway.


## References:
- https://micro.dev/getting-started
- https://github.com/micro/services
- https://github.com/xpunch/go-micro-example/blob/main/v4/README.md
- https://www.twilio.com/en-us/blog/a-practical-guide-to-creating-microservices-with-go-micro