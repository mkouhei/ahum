#!/usr/bin/make -f
# -*- makefile -*-

BIN := ahum
SRC := *.go
GOPKG := github.com/mkouhei/ahum
GOPATH := $(CURDIR)/_build:$(GOPATH)
export GOPATH


#all: clean format test build
all: clean format build

prebuild:
	go get github.com/surge/mqtt
	install -d $(CURDIR)/_build/src/$(GOPKG)
	cp -a $(CURDIR)/*.go $(CURDIR)/_build/src/$(GOPKG)


build: prebuild
	go build -o _build/$(BIN)

build-only:
	go build -o _build/$(BIN)

clean:
	@rm -rf _build/


format:
	for src in $(SRC); do \
		gofmt $$src > $$src.tmp ;\
		goimports $$src.tmp > $$src.tmp2 ;\
		mv -f $$src.tmp2 $$src ;\
	done


test: prebuild
	go test
