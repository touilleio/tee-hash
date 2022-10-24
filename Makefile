
VERSION=v1.0.0
GOOS=linux
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
VERSION_MAJOR=$(shell echo $(VERSION) | cut -f1 -d.)
VERSION_MINOR=$(shell echo $(VERSION) | cut -f2 -d.)
BINARY_NAME=tee-hash

all: ensure build

ensure:
	GOOS=${GOOS} $(GOCMD) mod download

clean:
	$(GOCLEAN)

build:
	GOOS=${GOOS} $(GOBUILD) \
		-o ${BINARY_NAME} .

test:
	$(GOTEST) ./...

