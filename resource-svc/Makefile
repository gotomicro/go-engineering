APP_NAME:=resource-svc
APP_PATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SCRIPT_PATH:=$(APP_PATH)/scripts
COMPILE_OUT:=$(APP_PATH)/bin/$(APP_NAME)

go:export EGO_DEBUG=true
go:export EGO_LOG_EXTRA_KEYS=X-Ego-Uid
go:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml
install:export EGO_DEBUG=true
install:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml --job=install

initialize:export EGO_DEBUG=true
initialize:
	@cd $(APP_PATH) && go run main.go --config=config/local.toml --job=initialize


build:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making build app<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@chmod +x $(SCRIPT_PATH)/build/*.sh
	@cd $(APP_PATH) && $(SCRIPT_PATH)/build/gobuild.sh $(APP_NAME) $(COMPILE_OUT)






