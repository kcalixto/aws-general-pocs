.PHONY: build clean deploy test serverless-deploy npmci

GO_BUILD := env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w"

build:
	export GO111MODULE=on
	export CGO_ENABLED=1

	${GO_BUILD} -o bin/handler/main handler/main.go
	chmod +x bin/handler/main
	mv bin/handler/main bin/handler/bootstrap
	cd bin/handler && zip handler.zip bootstrap

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
