PACKAGES = $(shell go list ./src/...)

.PHONY: build doc fmt lint run test vet test-cover-html test-cover-func collect-cover-data AUTHORS

# Prepend our vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
export GO15VENDOREXPERIMENT=1
#GOPATH := ${PWD}/vendor:${GOPATH}
# export GOPATH

# Used to populate version variable in main package.
VERSION=v1.0.0
BUILD_TIME=$(shell date -u +%Y-%m-%d:%H-%M-%S)
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD_TIME}"

# set image name and tags
REGISTRY=harbor.huayun.com/library/
IMAGE_NAME=hubimage
TAGS=${VERSION}
NAME=${REGISTRY}${IMAGE_NAME}:${TAGS}

# set go build opts
BUILD_OPTS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

default: build

build: build-base
	docker build -t ${NAME} .

build-base: fmt
	@echo "build $@"
	 ${BUILD_OPTS} go build ${LDFLAGS} -v -o ./bin/hubimage ./src/

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@echo "fmt $@"
	go fmt ./src/...

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	@echo "lint  $@"
	@test -z "$$(golint ./src/... | tee /dev/stderr)"

run: build
	@echo "run $@"
	./bin/hubimage

test:
	@echo "test $@"
	go test -cover=true ./src/...

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	go vet ./src/...


clean:
	@echo "clean $@"
	rm -rf ./bin

collect-cover-data:
	@echo "collect-cover-data $@"
	@echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -v -coverprofile=coverage.out -covermode=count $(pkg) || exit $$?;\
		if [ -f coverage.out ]; then\
		    tail -n +2 coverage.out >> coverage-all.out;\
                fi\
		;)

test-cover-html: test-cover-func
	@echo "test-cover-html $@"
	go tool cover -html=coverage-all.out -o coverage.html

test-cover-func:
	@echo "test-cover-func $@"
	go tool cover -func=coverage-all.out

AUTHORS: .mailmap .git/HEAD
	@echo " AUTHORS $@"
	git log --format='%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf > $@

