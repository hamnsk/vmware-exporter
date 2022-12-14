SERVICE_NAME := "vmware-exporter"
CURRENT_DIR = $(shell pwd)
GOPATH = $(shell echo ${HOME})
RELEASE ?=devel
COMMIT := git-$(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

GOOS?=linux
GOARCH?=amd64

.SILENT:

deps:
	go mod download

clean:
	rm -rf ./.bin/${SERVICE_NAME}

build: clean deps
	GOOS=${GOOS} GOARCH=${GOARCH} go build -mod=readonly \
	-ldflags "-s -w -X 'vmware-exporter/internal/version.Version=${RELEASE}' \
	-X 'vmware-exporter/internal/version.Commit=${COMMIT}' \
	-X 'vmware-exporter/internal/version.BuildTime=${BUILD_TIME}'" \
	-o ./.bin/${SERVICE_NAME}${VERSION} ./cmd/${SERVICE_NAME}/main.go

debug_build: clean deps
	GOOS=${GOOS} GOARCH=${GOARCH} go build \
	-ldflags "-s -w -X 'vmware-exporter/internal/version.Version=${RELEASE}' \
	-X 'vmware-exporter/internal/version.Commit=${COMMIT}' \
	-X 'vmware-exporter/internal/version.BuildTime=${BUILD_TIME}'" \
	-gcflags="all=-N -l" -o ./.bin/${SERVICE_NAME}${VERSION} ./cmd/${SERVICE_NAME}/main.go \
