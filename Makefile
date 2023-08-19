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

dependencies:
	@go get -u github.com/onsi/ginkgo/ginkgo
	@go get -u github.com/onsi/gomega/...

deps:dependencies
	@go mod tidy

cover:deps
	@go test ./... -coverprofile=c.out.tmp -coverpkg=./... && cat c.out.tmp | grep -v "_mock.go" > c.out

report:cover
	@go tool cover -func c.out | grep "total"

html-report:cover
	@go tool cover -html c.out

test:deps
	@go test ./...

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


.PHONY: cover