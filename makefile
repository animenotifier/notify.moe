# Makefile for Anime Notifier

GOCMD=@go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
BUILDJOBS=@./jobs/build.sh

server:
	$(GOBUILD)
jobs:
	$(BUILDJOBS)
install:
	$(GOINSTALL)
test:
	$(GOTEST)
bench:
	$(GOTEST) -bench .
depslist:
	$(GOCMD) list -f {{.Deps}} | sed -e 's/\[//g' -e 's/\]//g' | tr " " "\n"
all: server jobs

.PHONY: jobs