_: build deploy

build:
	export GO111MODULE=on
	export CGO_ENABLED=1

	env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/bootstrap handler/main.go
	chmod +x bin/bootstrap
	zip -j bin/main.zip bin/bootstrap

deploy: build
	sls deploy
