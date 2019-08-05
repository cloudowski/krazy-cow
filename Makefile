SHELL = /bin/bash
NAME ?= cloudowski/krazy-cow
SHORTNAME = krazy-cow

VERSION = $(shell git tag -l --points-at HEAD)
GITCOMMIT = $(shell git rev-list -1 HEAD --abbrev-commit)
ifeq ($(VERSION),)
	VERSION=latest
endif

BASEDIR = $(shell pwd)

all: clean test build

.PHONY: build buildimg buildimgtiny run runfg clean push kill test deploy
default: build

build:
	go build -ldflags="-w -s -X main.version=$(VERSION) -X main.gitCommit=$(GITCOMMIT)" -o cow *.go

buildimg: 
	docker build --build-arg VERSION=$(VERSION) --build-arg GITCOMMIT=$(GITCOMMIT) -t $(NAME):$(VERSION) -f Dockerfile .

buildimgtiny: 
	docker build  --build-arg VERSION=$(VERSION) --build-arg GITCOMMIT=$(GITCOMMIT) -t $(NAME):$(VERSION) -f Dockerfile.slim .

test: 
	go test -cover ./...


push: 
	docker push $(NAME):$(VERSION)

run: 
	docker run -p 8080:8080 --name=$(SHORTNAME) -d $(NAME):$(VERSION)

runfg: 
	docker run --rm -p 8080:8080 -ti $(NAME):$(VERSION)

kill:
	-docker rm -f $(SHORTNAME)

clean:
	-docker rm -f $(NAME)
	-docker rmi $(NAME):$(VERSION)
	cd "$(BASEDIR)/deploy" && kubectl delete -f .
	cd "$(BASEDIR)/deploy" && kubectl delete -f redis/ephemeral

deploy:
	cd "$(BASEDIR)/deploy" && kubectl apply -f .
	cd "$(BASEDIR)/deploy" && kubectl apply -f redis/ephemeral

getversion:
	@echo -n "$(VERSION)"