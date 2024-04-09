VERSION=$(shell git describe --tags)
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_COMMIT_DATE=$(shell git log -n1 --pretty='format:%cd' --date=format:'%Y%m%d')
REPO=github.com/bnb-chain/blob-syncer
IMAGE_NAME=ghcr.io/bnb-chain/blob-syncer

ldflags = -X $(REPO)/version.AppVersion=$(VERSION) \
          -X $(REPO)/version.GitCommit=$(GIT_COMMIT) \
          -X $(REPO)/version.GitCommitDate=$(GIT_COMMIT_DATE)

build_syncer:
ifeq ($(OS),Windows_NT)
	go build -o build/blob-syncer.exe -ldflags="$(ldflags)" cmd/blob-syncer/main.go
else
	go build -o build/blob-syncer -ldflags="$(ldflags)" cmd/blob-syncer/main.go
endif

build_server:
ifeq ($(OS),Windows_NT)
	go build -o build/blob-syncer-server.exe -ldflags="$(ldflags)" cmd/blob-syncer-server/main.go
else
	go build -o build/blob-syncer-server -ldflags="$(ldflags)" cmd/blob-syncer-server/main.go
endif

build:
	make build_syncer
	make build_server

install:
	go install cmd/blob-syncer/main.go
	go install cmd/blob-syncer-server/main.go

build_docker:
	docker build . -t ${IMAGE_NAME}

.PHONY: build install build_docker


###############################################################################
###                                Linting                                  ###
###############################################################################

golangci_lint_cmd=golangci-lint
golangci_version=v1.51.2

lint:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --timeout=10m

lint-fix:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0

format:
	bash scripts/format.sh

.PHONY: lint lint-fix format

swagger-gen:
	swagger generate server -f ./swagger.yaml -A blob-syncer --default-scheme=http


