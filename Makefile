MAKE = make

build:
	$(MAKE) build -C src/mymath
	$(MAKE) build -C src/mathapp

run:
	bin/mathapp
