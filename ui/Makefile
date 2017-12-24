PWD := $(shell pwd)
APP := spongebob-ui

DEV_UI_IMAGE := ckeyer/dev:node

VERSION := $(shell cat VERSION.txt)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

dev:
	docker run --rm -it \
	 --name $(APP)-dev \
	 -p 8080:8080 \
	 -v $(PWD):/opt/$(APP) \
	 -w /opt/$(APP) \
	 $(DEV_UI_IMAGE) bash
