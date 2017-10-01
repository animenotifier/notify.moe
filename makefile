# Makefile for Anime Notifier

GOCMD=@go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
BUILDJOBS=@./jobs/build.sh
BUILDPATCHES=@./patches/build.sh
BUILDBOTS=@./bots/build.sh
TSCMD=@tsc
IPTABLES=@sudo iptables

server:
	$(GOBUILD)
jobs:
	$(BUILDJOBS)
bots:
	$(BUILDBOTS)
patches:
	$(BUILDPATCHES)
js:
	$(TSCMD)
install:
	$(GOINSTALL)
test:
	$(GOTEST) github.com/animenotifier/... -v
bench:
	$(GOTEST) -bench .
tools:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/aerogo/pack
	go get -u github.com/aerogo/run
	go install github.com/aerogo/pack
	go install github.com/aerogo/run
versions:
	@go version
	@asd --version
assets:
	$(TSCMD)
	@pack
depslist:
	$(GOCMD) list -f {{.Deps}} | sed -e 's/\[//g' -e 's/\]//g' | tr " " "\n"
clean:
	find . -type f | xargs file | grep "ELF.*executable" | awk -F: '{print $1}' | xargs rm
ports:
	$(IPTABLES) -t nat -A OUTPUT -o lo -p tcp --dport 80 -j REDIRECT --to-port 4000
	$(IPTABLES) -t nat -A OUTPUT -o lo -p tcp --dport 443 -j REDIRECT --to-port 4001
all: assets server bots jobs patches

.PHONY: bots jobs patches ports
