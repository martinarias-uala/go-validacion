all: cover

LAMBDA_NAME := shapes-gin-api

build:
	mkdir -p .build
	cd .build && \
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap ../cmd/main.go


zip:build
	cd .build && \
	zip bootstrap.zip bootstrap

update:zip
	cd .build && \
	aws lambda update-function-code --function-name ${LAMBDA_NAME} --zip-file fileb://bootstrap.zip --region us-east-1 --profile $(profile)

deps:dependencies
	@go mod tidy

cover:deps
	$(HOME)/go/bin/ginkgo -r --progress --failOnPending --trace -coverpkg=./... -coverprofile=coverage.out -outputdir=./test

report:cover
	@go tool cover -html=./test/coverage.out -o ./test/coverage.html

lint-prepare:
	 brew list golangci-lint || brew install golangci-lint

lint:lint-prepare
	@golangci-lint run

clean:
	@rm -fr **/*.{out,xml,html}

clean-cache:
	@go clean -cache
	@go clean -testcache
	@go clean -modcache

dependencies:
	@go get -u github.com/onsi/ginkgo/ginkgo
	@go get -u github.com/onsi/gomega/...

test:deps
	@go test ./...

.PHONY: cover