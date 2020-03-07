GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-gateway cmd/main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t ccr.ccs.tencentyun.com/rainjoe/api-gateway:0.1.4
	docker login -u $(ccruser) -p $(ccrpassword) ccr.ccs.tencentyun.com
	docker push ccr.ccs.tencentyun.com/rainjoe/api-gateway:0.1.4

.PHONY: clean
clean:
	rm -f api-gateway