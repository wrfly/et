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

GLIDE_URL := https://github.com/Masterminds/glide/releases\
	/download/v0.13.1/glide-v0.13.1-linux-amd64.tar.gz

.PHONY: prepare
prepare:
ifeq (, $(shell which glide))
	wget $(GLIDE_URL) -O glide.tgz -q
	tar xzf glide.tgz
	linux-amd64/glide install
else
	@echo "Glide is installed"
endif

.PHONY: bin
bin:
	mkdir -p $(BIN_PATH)

.PHONY: build
build:
	go build $(GO_LDFLAGS) -o $(BIN) $(PKG)

.PHONY: test
test:
	go test -v --cover `glide nv`

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

.PHONY: push-dev-img
push-dev-img:
	docker push $(IMAGE):dev

.PHONY: tools
tools:
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/elazarl/go-bindata-assetfs/...

ASSET_FILE := server/asset/asset.go
.PHONY: asset
asset:
	go-bindata-assetfs -nometadata -prefix html -pkg asset html/...
	mv bindata_assetfs.go $(ASSET_FILE)
	gofmt -w $(ASSET_FILE)
