PWD := $(shell pwd)
APP := spongebob
PKG := github.com/ckeyer/$(APP)
GO := CGO_ENABLED=0 GOBIN=${PWD}/bundles go

VERSION := $(shell cat VERSION.txt)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

DEV_IMAGE := ckeyer/dev:go
DEV_UI_IMAGE := ckeyer/dev:node
IMAGE_NAME := $(if $(REGISTRY_URL), $(REGISTRY_URL)/chuanjian/$(APP):$(GIT_BRANCH), chuanjian/$(APP):$(GIT_BRANCH))

LD_FLAGS := -X $(PKG)/version.version=$(VERSION) -X $(PKG)/version.gitCommit=$(GIT_COMMIT) -w

HASH := $(shell which sha1sum || which shasum)

init:
	echo $(IMAGE_NAME)

local: generate
	$(GO) install -a -ldflags="$(LD_FLAGS)" ./cmd/spongebob

generate:
	protoc -I. \
	 -I/usr/local/include \
	 -I${GOPATH}/src \
	 -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	 --grpc-gateway_out=logtostderr=true:. \
	 --go_out=plugins=grpc:. protos/*.proto

build:
	docker run --rm \
	 --name $(APP)-building \
	 -e CGO_ENABLED=0 \
	 -v $(PWD):/opt/gopath/src/$(PKG) \
	 -w /opt/gopath/src/$(PKG) \
	 $(DEV_IMAGE) make local

run:
	docker run --rm \
	 --name $(APP)-dev-running \
	 -e CGO_ENABLED=0 \
	 -p 8089:8080 \
	 -v $(PWD):/opt/gopath/src/$(PKG) \
	 -w /opt/gopath/src/$(PKG) \
	 $(DEV_IMAGE) go run ./cmd/spond/main.go

image: build
	docker build -t $(IMAGE_NAME) .

test:
	go test -ldflags="$(LD_FLAGS)" $$(go list ./... |grep -v "vendor")

dev: dev-server

dev-server:
	docker run --rm -it \
	 --name $(APP)-dev \
	 -p 8089:8080 \
	 -v $(PWD):/opt/gopath/src/$(PKG) \
	 -w /opt/gopath/src/$(PKG) \
	 $(DEV_IMAGE) bash

dev-client:
	docker run --rm -it \
	 --name $(APP)-dev-client \
	 -v /var/run/docker.sock:/var/run/docker.sock \
	 -v $(PWD):/opt/gopath/src/$(PKG) \
	 -w /opt/gopath/src/$(PKG) \
	 $(DEV_IMAGE) bash
