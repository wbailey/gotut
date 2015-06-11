GOROOT := $(shell go env GOROOT)
GOPATH := $(shell go env GOPATH)
GOBUILD=go install
DEBUG=-a -v -work -x

.SILENT:

export

all: build run

debug: GOBUILD += $(DEBUG)
debug: all

build:
	$(MAKE) build -C src/mymath
	$(MAKE) build -C src/mystr
	$(MAKE) build -C src/mathapp

run:
	bin/mathapp
