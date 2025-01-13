
GOPATH:=$(shell go env GOPATH)

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/notes.proto

.PHONY: init
init:
	go go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/notes.proto

.PHONY: build
build:
	go build -o notes *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t notes:latest