MAKEPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
NAME:=rhinoctl

build: cleanup test
	cd $(MAKEPATH); go build -o $(NAME)

cleanup:
	cd $(MAKEPATH); go fmt ./...
	cd $(MAKEPATH); go vet ./...

test:
	cd $(MAKEPATH); go test ./...

cleanup-full: cleanup
	cd $(MAKEPATH); go mod tidy

docker: build
	cd $(MAKEPATH); docker build -t $(NAME)-test .

