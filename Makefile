GOROOT := $(go env GOROOT)
GOPATH := $(shell pwd)

export

build:
	$(MAKE) build -C src/mymath
	$(MAKE) build -C src/mystr
	$(MAKE) build -C src/mathapp

run:
	bin/mathapp
