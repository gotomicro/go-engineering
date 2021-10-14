APP_NAME:=svc-order
APP_PATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SCRIPT_PATH:=$(APP_PATH)/scripts
COMPILE_OUT:=$(APP_PATH)/bin/$(APP_NAME)
PROTO:=protos

go:export EGO_DEBUG=true
go:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml

build:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making build app<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@chmod +x $(SCRIPT_PATH)/build/*.sh
	@cd $(APP_PATH) && $(SCRIPT_PATH)/build/gobuild.sh $(APP_NAME) $(COMPILE_OUT)
	@echo -e "\n"


link:export PROTO_DIR=../chapter_proto
link:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>make $@<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@ if [[ ! -L $(APP_PATH)/$(PROTO) ]]; then ln -s $(PROTO_DIR) $(APP_PATH)/$(PROTO); echo "link created"; else echo "link exists"; fi;

gen-proto:
	@protoc -I ./protos --go_out=paths=source_relative:./$(PROTO) ./$(PROTO)/helloworld/*.proto
	@protoc -I ./protos --go-grpc_out=paths=source_relative:./$(PROTO) ./$(PROTO)/helloworld/*.proto
	@protoc -I ./protos --go-errors_out=paths=source_relative:./$(PROTO) ./$(PROTO)/helloworld/*.proto

gen-test:
	@protoc --proto_path=${APP_PATH}/${PROTO}  --go-test_out=out=./server/router,paths=source_relative:. ${APP_PATH}/${PROTO}/helloworld/*.proto
