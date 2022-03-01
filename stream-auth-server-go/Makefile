.DEFAULT_GOAL := help

add_local_lib: ##add_local_lib
	@echo "add_local_lib"
	@go mod edit -replace github.com/gitJaesik/stream_go_srvs/streamgolib=../streamgolib

generate_statik: ## generate_statik
	@statik -m -f -src public/

env: ## set env for local lib
	@echo "set env for local lib"
	@go env -w GOPRIVATE=github.com/*
	@go env -w GOINSECURE=github.com/*

install: ## install go dependencies
	@echo "Install go dependencies"
	@go get -u \
		github.com/spf13/viper \
		github.com/spf13/pflag \
		go.uber.org/zap \
		github.com/lestrrat-go/file-rotatelogs \
		github.com/dgrijalva/jwt-go \
		golang.org/x/crypto/bcrypt \
		go.mongodb.org/mongo-driver/mongo \
		github.com/fatih/color \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/rakyll/statik \
		github.com/bufbuild/buf/cmd/buf \
		github.com/google/wire/cmd/wire \
		github.com/go-audio/wav \
		gopkg.in/confluentinc/confluent-kafka-go.v1/kafka \
		golang.org/x/oauth2 \
		cloud.google.com/go/compute/metadata \
		github.com/go-playground/validator \
		github.com/stretchr/testify \
		github.com/google/uuid \
		github.com/cespare/reflex \
		github.com/onsi/ginkgo/v2 \
		github.com/go-co-op/gocron \
		google.golang.org/grpc

tool: ## install go tool
	@echo "install go tool"
	@go get -u golang.org/x/tools/cmd/stringer
	@go get -u github.com/pilu/fresh

generate: ## protobuf compile and make static assets
	@echo "protobuf compile and make static assets"
	@buf generate --template buf.gen.yaml .

	# Generate static assets for OpenAPI UI
	# @statik -m -f -src third_party/OpenAPI/

start: ## start server
	@echo "start server"
	@./run_server.sh

tag: ## make tag
	@echo "make tag"
	@./scripts/auto_increase_tag.sh
	# @/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/gitJaesik/shell_script_repo/master/auto_increase_tag.sh)"

build_linux: ## build_linux server
	@echo "build linux server"
	@./run_build_linux.sh

build_local: ## build_local server
	@echo "build local server"
	@./run_build_local.sh

start_all: ## start_all
	@echo "start_all"
	@scripts/start_all.sh

stop_all: ## start_all
	@echo "stop_all"
	@scripts/stop_all.sh

start_mongo: ## start_mongo
	@echo "start_mongo"
	@scripts/start_mongo.sh

stop_mongo: ## stop_mongo
	@echo "stop_mongo"
	@scripts/stop_mongo.sh

start_kafka: ## start kafka docker
	@echo "start kafka docker"
	@scripts/start_kafka.sh

stop_kafka: ## stop kafka docker
	@echo "stop kafka docker"
	@scripts/stop_kafka.sh

test: ## ginkgo -r
	@echo "ginkgo -r"
	@ginkgo -r

reflex: ## run reflex to watch and run
	@echo "run reflex to watch and run"
	@reflex -s -g reflex.conf -- reflex -c reflex.conf

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html


