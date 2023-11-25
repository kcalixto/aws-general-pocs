_: build deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main handler/main.go

deploy: build
	sls deploy