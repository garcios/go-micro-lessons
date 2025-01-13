
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/stock.proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/stock.proto

.PHONY: build
build:
	go build -o stock *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	podman build . -t stock:latest