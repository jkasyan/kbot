APP=kbot
REPOSITORY=ghcr.io
USERNAME=jkasyan # TODO: env var ???
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETARCH=$(shell dpkg --print-architecture)
IMAGE_ID=${VERSION}-${TARGETARCH}
IMAGE_TAG=${REPOSITORY}/${USERNAME}/${APP}:${IMAGE_ID}

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

get:
	go get

build: format get
	@echo "OS ---> ${OS}"
	CGO_ENABLED=0 GOOS=${OS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags "-X="github.com/JKasyan/cmd.appVersion=${VERSION}

linux:
	$(MAKE) build OS=linux

arm:
	$(MAKE) build OS=arm

macos:
	$(MAKE) build OS=macos

windows:
	$(MAKE) build OS=windows

# TODO: image tag ??? v1.0.4-47c6f54-7cbdc75-ce55add-amd64
image:
	@echo "tag: ${IMAGE_TAG}"
	docker build -t ${IMAGE_TAG} . 

push:
    @echo "tag: ${IMAGE_TAG}"
	docker push ${IMAGE_TAG}

clean:
	@echo "remove image with tag: ${IMAGE_ID}"
	docker rmi $(shell docker images | grep ${IMAGE_ID} | awk '{print $$3}')