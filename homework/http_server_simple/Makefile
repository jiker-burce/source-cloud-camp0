export tag=v1.0.1
root:
	export ROOT=github.com/bruce-gu/golang

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t storefeegmail/httpserver:${tag} .

push: release
	echo "pushing storefeegmail/httpserver"
	docker push storefeegmail/httpserver:${tag}
