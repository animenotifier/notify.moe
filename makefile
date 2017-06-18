# Makefile for Anime Notifier

GOCMD=@go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
BUILDJOBS=@./jobs/build.sh
BUILDPATCHES=@./patches/build.sh
IPTABLES=@sudo iptables

server:
	$(GOBUILD)
jobs:
	$(BUILDJOBS)
patches:
	$(BUILDPATCHES)
install:
	$(GOINSTALL)
test:
	$(GOTEST)
bench:
	$(GOTEST) -bench .
versions:
	@go version
	@asd --version
assets:
	@pack
depslist:
	$(GOCMD) list -f {{.Deps}} | sed -e 's/\[//g' -e 's/\]//g' | tr " " "\n"
clean:
	find . -type f | xargs file | grep "ELF.*executable" | awk -F: '{print $1}' | xargs rm
ports:
	$(IPTABLES) -t nat -A OUTPUT -o lo -p tcp --dport 80 -j REDIRECT --to-port 4000
	$(IPTABLES) -t nat -A OUTPUT -o lo -p tcp --dport 443 -j REDIRECT --to-port 4001
all: assets server jobs patches

.PHONY: jobs patches ports
