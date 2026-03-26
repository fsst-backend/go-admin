PROJECT:=lotus-go-admin
SHELL = /bin/bash

BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
VERSION=git-$(subst /,-,$(BRANCH))-$(shell git describe --tags --always --dirty)
IMAGE_TAG=$(VERSION)
IMAGE_REPO=docker.fastdocker.com:5000
PKG=lotus/go-admin

PROTO_DIR := grpc/proto
PB_DIR := grpc/pb

PROTO_FILES := $(shell find $(PROTO_DIR) -name "*.proto")


.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -o ./build/mac/go-admin .

# make build-linux
build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0  \
		go build -v -tags netgo -installsuffix netgo -installsuffix cgo -ldflags '-w -s' -ldflags "-X main.Version=$(VERSION)" \
		-o ./build/linux/go-admin .
	@echo "build successful"

.PHONY: print-tag
print-tag:
	@echo $(IMAGE_TAG)

push: build-linux
	docker build --platform linux/amd64 -t ${IMAGE_REPO}/${PROJECT}:${IMAGE_TAG} .
	docker push ${IMAGE_REPO}/${PROJECT}:${IMAGE_TAG}
	docker rmi ${IMAGE_REPO}/${PROJECT}:${IMAGE_TAG}
	docker image prune -f


.PHONY: proto
proto:
	@echo "==> Generating protobuf files..."
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(PB_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PB_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)
	@echo "==> Done."

# 将 violet 的 backend-api.json 复制到 docs/violet/
.PHONY: copy-violet-api
copy-violet-api:
	@mkdir -p docs/violet
	@cp ../violet/doc/backend-api.json docs/violet/backend-api.json
	@echo "==> copied ../violet/doc/backend-api.json -> docs/violet/backend-api.json"