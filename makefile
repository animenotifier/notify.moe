# Makefile for Anime Notifier

GOCMD=@go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOTEST=@./utils/test/go-test-color.sh
BUILDJOBS=@./jobs/build.sh
BUILDPATCHES=@./patches/build.sh
BUILDBOTS=@./bots/build.sh
TSCMD=@tsc
IPTABLES=@sudo iptables
IP6TABLES=@sudo ip6tables
PACK:=$(shell command -v pack 2> /dev/null)
RUN:=$(shell command -v run 2> /dev/null)
GOIMPORTS:=$(shell command -v goimports 2> /dev/null)

# Determine the name of the platform
OSNAME=

ifeq ($(OS),Windows_NT)
	OSNAME = WINDOWS
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OSNAME = LINUX
	endif
	ifeq ($(UNAME_S),Darwin)
		OSNAME = OSX
	endif
endif

# Build targets
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
	$(GOTEST) github.com/animenotifier/notify.moe -v -cover
bench:
	$(GOTEST) -bench .
pack:
	go get -u github.com/aerogo/pack
	go install github.com/aerogo/pack
run:
	go get -u github.com/aerogo/run
	go install github.com/aerogo/run
goimports:
	go get -u golang.org/x/tools/cmd/goimports
	go install golang.org/x/tools/cmd/goimports
tools:
ifeq ($(OSNAME),OSX)
	brew install coreutils
endif
ifndef GOIMPORTS
	@make goimports
endif
ifndef PACK
	@make pack
endif
ifndef RUN
	@make run
endif
versions:
	@go version
	$(TSCMD) --version
assets:
	$(TSCMD)
	@pack
deps:
	# Ignore errors using the "-" because components directory can not be fetched.
	@-go get -t -v ./...
depslist:
	$(GOCMD) list -f {{.Deps}} | sed -e 's/\[//g' -e 's/\]//g' | tr " " "\n"
clean:
	find . -type f | xargs file | grep "ELF.*executable" | awk -F: '{print $1}' | xargs rm
ports:
ifeq ($(OSNAME),LINUX)
	$(IPTABLES) -t nat -A OUTPUT -o lo -p tcp --dport 80 -j REDIRECT --to-port 4000
	$(IPTABLES) -t nat -A OUTPUT -o lo -p tcp --dport 443 -j REDIRECT --to-port 4001
	$(IP6TABLES) -t nat -A OUTPUT -o lo -p tcp --dport 80 -j REDIRECT --to-port 4000
	$(IP6TABLES) -t nat -A OUTPUT -o lo -p tcp --dport 443 -j REDIRECT --to-port 4001
endif
ifeq ($(OSNAME),OSX)
	@echo "rdr pass inet proto tcp from any to any port 443 -> 127.0.0.1 port 4001" | sudo pfctl -ef -
endif
browser:
ifeq ($(OSNAME),LINUX)
	@google-chrome --ignore-certificate-errors
endif
ifeq ($(OSNAME),OSX)
	@/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --ignore-certificate-errors
endif
all: tools assets server bots jobs patches

.PHONY: tools assets server bots jobs patches ports clean versions
