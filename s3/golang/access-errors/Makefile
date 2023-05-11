run: clean build deploy

clean:
	rm -rf handler/main

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o handler/main handler/main.go

deploy:
	serverless deploy --verbose