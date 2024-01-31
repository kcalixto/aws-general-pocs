build:
	export GO111MODULE=on
	export CGO_ENABLED=0

	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/handler handler/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/trigger trigger/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
