PWD := $(shell pwd)
APP := spongebob
PKG := github.com/ckeyer/$(APP)
GO := CGO_ENABLED=0 GOBIN=${PWD}/bundles go
HASH := $(shell which sha1sum || which shasum)

VERSION := $(shell cat VERSION.txt)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

DEV_IMAGE := ckeyer/dev:go
DEV_UI_IMAGE := ckeyer/dev:node
IMAGE_NAME := ckeyer/$(APP):$(GIT_BRANCH)

LD_FLAGS := -X $(PKG)/pkgs/version.version=$(VERSION) -X $(PKG)/pkgs/version.gitCommit=$(GIT_COMMIT) -w

init:
	echo $(IMAGE_NAME)

local: generate
	$(GO) install -a -ldflags="$(LD_FLAGS)" .

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
	 $(DEV_IMAGE) ${GO} run ./main.go

image: build
	docker build -t $(IMAGE_NAME) .

push:
	docker push $(IMAGE_NAME)

test:
	${GO} test -ldflags="$(LD_FLAGS)" $$(go list ./... |grep -v "vendor")

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
