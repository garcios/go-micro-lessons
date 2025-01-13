
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --openapi_out=./config --proto_path=. --micro_out=. --go_out=. proto/helloworld.proto

.PHONY: api
api:
	protoc --openapi_out=./config --proto_path=. proto/helloworld.proto

.PHONY: build
build: proto

	go build -o helloworld-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t helloworld-srv:latest