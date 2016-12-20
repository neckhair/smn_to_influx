SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=smn_to_influx
BUILD=$(shell git rev-parse HEAD)
VERSION=$(shell cat VERSION)
LDFLAGS=-ldflags "-X github.com/neckhair/smn_to_influx/core.Version=${VERSION} -X github.com/neckhair/smn_to_influx/core.Build=${BUILD}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	GOOS=linux go build ${LDFLAGS} -o ${BINARY} main.go

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clear
clear:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
