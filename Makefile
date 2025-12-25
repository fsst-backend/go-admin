PROJECT:=lotus-go-admin
SHELL = /bin/bash

BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
VERSION=git-$(subst /,-,$(BRANCH))-$(shell git describe --tag --dirty)
IMAGE_TAG=$(VERSION)
IMAGE_REPO=docker.fastdocker.com:5000
PKG=lotus/go-admin

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -o ./build/mac/go-admin .

# make build-linux
build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0  \
		go build -v -tags netgo -installsuffix netgo -installsuffix cgo -ldflags '-w -s' -ldflags "-X main.Version=$(VERSION)" \
		-o ./build/linux/go-admin .
	@echo "build successful"


push: build-linux
	docker build -t ${IMAGE_REPO}/${PROJECT}:${IMAGE_TAG} .
	docker push ${IMAGE_REPO}/${PROJECT}:${IMAGE_TAG}
	docker rmi ${IMAGE_REPO}/${PROJECT}:${IMAGE_TAG}
	docker image prune -f