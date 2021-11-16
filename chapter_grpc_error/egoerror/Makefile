APP_NAME:=egoerror
APP_PATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SCRIPT_PATH:=$(APP_PATH)/scripts
COMPILE_OUT:=$(APP_PATH)/bin/$(APP_NAME)

cli:export EGO_DEBUG=true
cli:
	@go run client/main.go --config=client/config.toml

svc:export EGO_DEBUG=true
svc:
	@go run server/main.go --config=server/config.toml

gen:
	@protoc -I ./helloworld --go_out=paths=source_relative:./helloworld ./helloworld/*.proto
	@protoc -I ./helloworld --go-grpc_out=paths=source_relative:./helloworld ./helloworld/*.proto
	@protoc -I ./helloworld --go-errors_out=paths=source_relative:./helloworld ./helloworld/*.proto
