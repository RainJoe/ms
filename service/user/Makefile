GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc -I proto/user/ proto/user/user.proto --go_out=plugins=grpc:proto/user

.PHONY: build
build: proto

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-srv cmd/main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t ccr.ccs.tencentyun.com/rainjoe/user-srv:0.1.15
	docker login -u $(ccruser) -p $(ccrpassword) ccr.ccs.tencentyun.com
	docker push ccr.ccs.tencentyun.com/rainjoe/user-srv:0.1.15

.PHONY: clean
clean:
	rm -rf user-srv