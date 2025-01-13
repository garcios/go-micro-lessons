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

Dashboard
```shell
go install github.com/go-micro/dashboard@latest
```


## mod file
```go
module github.com/oskiegarcia/go-micro-lessons

go 1.22.2

require (
	github.com/golang/protobuf v1.5.4
	google.golang.org/protobuf v1.36.2
	micro.dev/v4 v4.6.0
)
```



### References:
- https://micro.dev/getting-started
- https://github.com/micro/services
- https://github.com/xpunch/go-micro-example/blob/main/v4/README.md