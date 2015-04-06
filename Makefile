SELFPKG := bitbucket.org/agurha/tunnel
VERSION := 0.1
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
PKGS := \
build \
PKGS := $(addprefix bitbucket.org/agurha/tunnel/pkg/,$(PKGS))
.PHONY := default deps vendor build build-dist godep rice

BUILDTAGS=debug

default: all


all: embedClient embedServer build



build:
	go build -o bin/tunnel-cli -tags '$(BUILDTAGS)' -ldflags "-X main.version $(VERSION)dev-$(SHA)" $(SELFPKG)/client/main
	go build -o bin/tunnel-server -tags '$(BUILDTAGS)' -ldflags "-X main.version $(VERSION)dev-$(SHA)" $(SELFPKG)/server/main

build-dist: godep
	godep go build -o bin/tunnel-cli -ldflags "-X main.version $(VERSION)-$(SHA)" $(SELFPKG)/client/main
	godep go build -o bin/tunnel-server -ldflags "-X main.version $(VERSION)-$(SHA)" $(SELFPKG)/server/main

bump-deps: deps vendor

deps:
	go get -u -t -v ./...

vendor: godep
	git submodule update --init --recursive
	godep save ./...

embedClient: clientTLS clientPublic rice
	cd client/main && rice embed

clientTLS:
	cd client/assets/tls

clientPublic:
	cd client/views/web/public && find *

embedServer: serverTLS rice
	cd server/main && rice embed

serverTLS:
	cd server/assets/tls

godep:
	go get github.com/tools/godep

rice:
	go install github.com/GeertJohan/go.rice/rice