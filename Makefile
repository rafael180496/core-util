PROJECT_NAME := "libcore"
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
VERSION := $(shell cat version)

# OS / Arch we will build our binaries for
OSARCH := "linux/amd64 linux/386 windows/amd64 windows/386 darwin/amd64 darwin/386"

.SHELLFLAGS = -c # Run commands in a -c flag
.SILENT: ; # no need for @
.ONESHELL: ; # recipes execute in same shell

default: build

all: test build

init: gen-mock

clean: ## Clean dev files
	go clean -i ./...
	rm -vf \
		"./${PROJECT_NAME}" \
		./coverage.* \
		./cpd.*

dependencies: dep ## Get ALL dependencies
	@go mod download

tidy:  ## Execute tidy comand
	@go mod tidy

build-out: tidy ## Build the binary file
	go build -i -v $(PKG_LIST) -o ./dist/$(PROJECT_NAME)_$(VERSION)

cross-build: tidy ## Build the app for multiple os/arch
	@gox -osarch=$(OSARCH) -output "dist/{{.OS}}_{{.Arch}}/${PROJECT_NAME}"

install: build ## Build the binary file
	@go install

package-release: cross-build ## Package release files
	mkdir -v -p $(CURDIR)/releases
	for release in $$(find ./dist -mindepth 1 -maxdepth 1 -type d 2>/dev/null); do \
		platform=$$(basename $$release); \
		tar -cvzf ./releases/${PROJECT_NAME}_${VERSION}_$$platform.tgz -C ./dist/$$platform .; \
	done

release: package-release ## Execute release
	ghr --replace $(VERSION) releases/

lint: ## Execute lint
	@golangci-lint run ./...

fmt: ## Formmat src code files
	@go fmt ${PKG_LIST}

cpd: ## CPD
	dupl -t 200 -html >cpd.html

test: ## Execute test
	@echo "go test ${PKG_LIST}"
	go test -i ${PKG_LIST} || exit 1
	echo ${PKG_LIST} | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

race: ## Run data race detector
	@go test -race -short ${PKG_LIST}

bench: ## Run benchmarks
	go test -bench ${PKG_LIST}

msan: ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

vet: ## Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string
	@echo "go vet ."
	@go vet ${PKG_LIST} ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

misspell: ## One way of improving the accuracy of your writing is to spell things right.
	@misspell -locale US  .

coverage: ## Generate global code coverage report
	./scripts/coverage.sh;

gen-mock: ## Execute mockgen generatio go mocks
	mockgen -destination ./app/mock/roundtripper.go -package mhttp net/http RoundTripper

security: ## Execute go sec security step
	gosec -tests ./...

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: \
	all \
	init \
	clean \
	dep \
	dependencies \
	tidy \
	build \
	build-out \
	cross-build \
	install \
	package-release \
	release \
	lint \
	fmt \
	cpd \
	race \
	bench \
	vet \
	misspell \
	coverage \
	gen-mock \
	security
