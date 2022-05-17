Go=go
App=clean-gin-server
type=prod

docker_registry=docker-${type}-local
image=${docker_registry}/${App}
version=v`git rev-list HEAD --count`-${type}

.PHONY: hello:
	@echo HELLO FROM MAKE!

.PHONY: bye:
	@echo MAKE IS EXITING...

.PHONY: test:
	@$(Go) test -cover -v ./...

.PHONY: run:
	@$(Go) run .

.PHONY: docker_build:
	@echo '-------------------------------------------'
	@echo $(image):$(version)
	@echo '-------------------------------------------'
	@docker build -t $(image):$(version) -t $(image):latest --no-cache .

.PHONY: docker_push:
#	@docker push $(image) --all
	@docker push $(image):$(version)
	@docker push $(image):latest

.PHONY: docker_rmi:
	@docker rmi $(image):$(version)
	@docker rmi $(image):latest

.PHONY: docker_job: docker_build docker_push docker_rmi

.PHONY: clean:
	@$(Go) clean

.PHONY: deploy: hello test build_linux clean

.PHONY: build_windows:
	@set CGO_ENABLED=0& set GOOS=windows& set GOARCH=amd64& $(Go) build -a -ldflags "-s" .
	@echo "WINDOWS BUILD COMPLETE"

.PHONY: build_linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(Go) build -race -a -ldflags="-s -w" .
	@echo "BUILD COMPLETE"

# CROSS COMPILE
.PHONY: cc:
	@echo "Cross Compile"
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 $(Go) build .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(Go) build .
	CGO_ENABLED=0 GOOS=windows GOARCH=arm $(Go) build .

.PHONY: cc_windows:
	@echo "Cross Compile Windows"
	set CGO_ENABLED=0& set GOOS=darwin& set GOARCH=386& go build -o bin/$(App)-386.exe .
	set CGO_ENABLED=0& set GOOS=linux& set GOARCH=amd64& $(Go) build -o bin/$(App)-linux-amd64.exe .
	set CGO_ENABLED=0& set GOOS=windows& set GOARCH=arm& $(Go) build -o bin/$(App)-windows-arm.exe .

.PHONY: all: hello docker_job bye