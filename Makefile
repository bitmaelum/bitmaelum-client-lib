BUILD_MODE?=c-shared
OUTPUT_DIR?=output
BINDING_NAME?=libbitmaelum_bridge
BINDING_FILE?=$(BINDING_NAME).so
BINDING_ARGS?=
BINDING_OUTPUT?=$(OUTPUT_DIR)/binding

default: fmt test

deps:
	go mod download

test:
	go test ./bridge/... -short -count 1

fmt:
	go fmt ./...

clean:
	rm -rf output

binding: deps
	mkdir -p $(BINDING_OUTPUT)
ifeq ($(BUILD_MODE), )
	go build -ldflags="-w -s" -o $(BINDING_OUTPUT)/$(BINDING_FILE) $(BINDING_ARGS) binding/main.go
else
	go build -ldflags="-w -s" -o $(BINDING_OUTPUT)/$(BINDING_FILE) -buildmode=$(BUILD_MODE) $(BINDING_ARGS) binding/main.go
endif

include Makefile.android
include Makefile.ios
include Makefile.darwin
#include Makefile.linux
#include Makefile.windows
#include Makefile.gomobile
#include Makefile.wasm
