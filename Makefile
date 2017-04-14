PWD := $(shell pwd)
APP := spongebob
PKG := github.com/ckeyer/$(APP)

VERSION := $(shell cat VERSION.txt)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

DEV_IMAGE := ckeyer/dev
DEV_UI_IMAGE := ckeyer/dev:node
IMAGE_NAME := $(if $(REGISTRY_URL), $(REGISTRY_URL)/chuanjian/$(APP):$(GIT_BRANCH), chuanjian/$(APP):$(GIT_BRANCH))

LD_FLAGS := -X $(PKG)/version.version=$(VERSION) -X $(PKG)/version.gitCommit=$(GIT_COMMIT) -w

HASH := $(shell which sha1sum || which shasum)

init:
	echo $(IMAGE_NAME)

local:
	go build -a -ldflags="$(LD_FLAGS)" -o bundles/spond ./cmd/spond
	$(HASH) bundles/spond
	go build -a -ldflags="$(LD_FLAGS)" -o bundles/sponctl ./cmd/sponctl
	$(HASH) bundles/sponctl

build:
	docker run --rm \
	 --name $(APP)-building \
	 -e CGO_ENABLED=0 \
	 -v $(PWD):/opt/gopath/src/$(PKG) \
	 -w /opt/gopath/src/$(PKG) \
	 $(DEV_IMAGE) make local

image: build
	docker build -t $(IMAGE_NAME) .

test:
	go test -ldflags="$(LD_FLAGS)" $$(go list ./... |grep -v "vendor")

dev:
	docker run --rm -it \
	 --name $(APP)-dev \
	 -p 8080:8080 \
	 -v /var/run/docker.sock:/var/run/docker.sock \
	 -v $(PWD):/opt/gopath/src/$(PKG) \
	 -w /opt/gopath/src/$(PKG) \
	 $(DEV_IMAGE) bash
