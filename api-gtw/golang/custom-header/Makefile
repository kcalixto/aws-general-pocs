.PHONY: build clean deploy test serverless-deploy npmci

GO_BUILD := env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w"

build:
	export GO111MODULE=on
	export CGO_ENABLED=1

	${GO_BUILD} -o bin/bootstrap handler/main.go
	chmod +x bin/bootstrap
	zip -j bin/api.zip bin/bootstrap

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
