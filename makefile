# Makefile for Anime Notifier

# Constants
GOTEST=@./utils/test/go-test-color.sh
GOBINARIES=`go env GOPATH`/bin
PACK=$(GOBINARIES)/pack

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

# builds the server executable
server:
	@go build -v

# installs development tools
tools:
ifeq ($(OSNAME),OSX)
	brew install coreutils
endif
	go install github.com/aerogo/pack/...
	go install github.com/aerogo/run/...
	go install github.com/itchyny/gojq/...

# compiles assets for the server
assets:
	@tsc
	@$(PACK)

# cleans all binaries and generated files
clean:
	find . -type f | xargs file | grep "ELF.*executable" | awk -F: '{print $1}' | xargs rm
	find . -type f | grep /scripts/ | grep .js | xargs rm
	rm -rf ./components

# forwards local ports 80 and 443 to 4000 and 4001
ports:
ifeq ($(OSNAME),LINUX)
	@sudo iptables -t nat -A OUTPUT -o lo -p tcp --dport 80 -j REDIRECT --to-port 4000
	@sudo iptables -t nat -A OUTPUT -o lo -p tcp --dport 443 -j REDIRECT --to-port 4001
	@sudo ip6tables -t nat -A OUTPUT -o lo -p tcp --dport 80 -j REDIRECT --to-port 4000
	@sudo ip6tables -t nat -A OUTPUT -o lo -p tcp --dport 443 -j REDIRECT --to-port 4001
endif
ifeq ($(OSNAME),OSX)
	@echo "rdr pass inet proto tcp from any to any port 443 -> 127.0.0.1 port 4001" | sudo pfctl -ef -
endif

# downloads the database
db:
	@./db/build.sh

# installs systemd service files for all required services
services:
	@./services/build.sh

# builds all background jobs
jobs:
	@./jobs/build.sh

# builds all bots
bots:
	@./bots/build.sh

# builds all patches
patches:
	@./patches/build.sh

test:
	$(GOTEST) github.com/animenotifier/notify.moe -v -cover

bench:
	$(GOTEST) -run=^$ -bench .

all: tools assets server bots jobs patches

.PHONY: tools assets server bots jobs patches services db ports clean
