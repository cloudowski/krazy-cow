SHELL := /bin/bash
NAME = cloudowski/trapped-cow
SHORTNAME = trapped-cow

# VERSION?=$(shell git tag -l --points-at HEAD)
VERSION?=latest

all: clean build

.PHONY: build tinybuild run runfg clean push kill
default: build

build: 
	docker build -t $(NAME):$(VERSION) .

tinybuild: 
	docker build -f Dockerfile.slim -t $(NAME):$(VERSION) .

push: tinybuild
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
