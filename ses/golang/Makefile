_: build deploy

clean:
	rm -rf bin

build: clean
	export GO111MODULE=on
	export CGO_ENABLED=1

	env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/bootstrap handler/main.go
	chmod +x bin/bootstrap
	zip -j bin/main.zip bin/bootstrap
	zip -r bin/main.zip resources/*

deploy:
	sls deploy