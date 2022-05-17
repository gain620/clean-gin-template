Go=go
App=clean-gin-server
type=prod

docker_registry=ghcr.io/${USERNAME}
image=${docker_registry}/${App}
version=v`git rev-list HEAD --count`-${type}

hello:
	@echo HELLO FROM MAKE!

bye:
	@echo MAKE IS EXITING...

test:
	@$(Go) test -cover -v ./...

run:
	@$(Go) run .

docker_build:
	@echo '-------------------------------------------'
	@echo $(image):$(version)
	@echo '-------------------------------------------'
	@docker build -t $(image):$(version) -t $(image):latest --no-cache .

docker_login:
	@echo ${GHCR_TOKEN} | docker login ghcr.io -u ${USERNAME} --password-stdin

docker_push:
#	@docker push $(image) --all
	@docker push $(image):$(version)
	@docker push $(image):latest

docker_rmi:
	@docker rmi $(image):$(version)
	@docker rmi $(image):latest

docker_job: docker_build docker_login docker_push docker_rmi

clean:
	@$(Go) clean

deploy: hello test build_linux clean

build_windows:
	@set CGO_ENABLED=0& set GOOS=windows& set GOARCH=amd64& $(Go) build -a -ldflags "-s" .
	@echo "WINDOWS BUILD COMPLETE"

build_linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(Go) build -race -a -ldflags="-s -w" .
	@echo "BUILD COMPLETE"

# CROSS COMPILE
cc:
	@echo "Cross Compile"
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 $(Go) build .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(Go) build .
	CGO_ENABLED=0 GOOS=windows GOARCH=arm $(Go) build .

cc_windows:
	@echo "Cross Compile Windows"
	set CGO_ENABLED=0& set GOOS=darwin& set GOARCH=386& go build -o bin/$(App)-386.exe .
	set CGO_ENABLED=0& set GOOS=linux& set GOARCH=amd64& $(Go) build -o bin/$(App)-linux-amd64.exe .
	set CGO_ENABLED=0& set GOOS=windows& set GOARCH=arm& $(Go) build -o bin/$(App)-windows-arm.exe .

.PHONY: all
all: hello docker_job bye