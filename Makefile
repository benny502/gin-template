GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
APP_MAIN_DIR=cmd/app
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
		   --go_out=paths=source_relative:./api \
		   --go-grpc_out=paths=source_relative:./api \
		   $(API_PROTO_FILES)

.PHONY: run
# run
run:
	cd $(APP_MAIN_DIR) && go run .

.PHONY: build
# 自动根据平台编译二进制文件
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# 生成应用所需的文件
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...


.PHONY: wire
# wire
wire:
	cd $(APP_MAIN_DIR) && wire

.PHONY: clean
# clean
clean:
	cd bin && rm -rf *

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help