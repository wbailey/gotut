GOROOT := $(shell go env GOROOT)
GOPATH := $(shell go env GOPATH)
GOBUILD= go install -a -v

export

all: build run

build:
	$(MAKE) build -C src/mymath
	$(MAKE) build -C src/mystr
	$(MAKE) build -C src/mathapp

run:
	bin/mathapp
