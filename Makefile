
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=paths=source_relative:. --go_out=paths=source_relative:. proto/helloworld/helloworld.proto

.PHONY: build
build: proto
	go build -o helloworld-service main.go plugin.go generate.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t helloworld-service:latest

.PHONY: tools
tools: ## 開発に必要なツールをインストールします
	go mod download
	go generate -tags=tools tools.go

.PHONY: help
help: __ ## ヘルプを表示します
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@cat $(MAKEFILE_LIST) \
	| grep -e "^[a-zA-Z_/\-]*: *.*## *" \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}'
