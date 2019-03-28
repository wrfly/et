NAME := et
AUTHOR := wrfly

BIN_PATH := bin
BIN := $(BIN_PATH)/$(NAME)
IMAGE := $(AUTHOR)/$(NAME)
PKG := github.com/$(AUTHOR)/$(NAME)/main

VERSION := $(shell cat VERSION)
COMMITID := $(shell git rev-parse --short HEAD)
BUILDAT := $(shell date +%Y-%m-%d)

CTIMEVAR := -X main.CommitID=$(COMMITID) \
        -X main.Version=$(VERSION) \
        -X main.BuildAt=$(BUILDAT)
GO_LDFLAGS := -ldflags "-s -w $(CTIMEVAR)" -tags netgo

export GO111MODULE=on

.PHONY: bin
bin:
	mkdir -p $(BIN_PATH)

.PHONY: build
build:
	go build $(GO_LDFLAGS) -o $(BIN) $(PKG)

.PHONY: test
test:
	go test -v --cover ./...

.PHONY: dev
dev: bin asset build
	@echo source config.env
	$(BIN)

.PHONY: image
image:
	docker build -t $(IMAGE):$(VERSION) -t $(IMAGE) -t $(IMAGE):dev .

.PHONY: push-img
push-img:
	docker push $(IMAGE)

.PHONY: push-tag
push-tag:
	docker push $(IMAGE)
	docker push $(IMAGE):$(VERSION)

.PHONY: push-img-dev
push-img-dev:
	docker push $(IMAGE):dev

.PHONY: tools
tools:
	go get github.com/wrfly/bindata

ASSET_FILE := server/asset/asset.go
.PHONY: asset
asset:
	bindata -pkg github.com/$(AUTHOR)/$(NAME)/server/asset \
		-resource html/
