#!/bin/bash

echo "Checking prerequisites..."

if hash go 2>/dev/null; then
	go version
else
	echo "Go is not installed"
	exit
fi

if hash tsc 2>/dev/null; then
	tsc --version
else
	echo "TypeScript is not installed"
	exit
fi

if hash git 2>/dev/null; then
	git version
else
	echo "Git is not installed"
	exit
fi

if hash git-lfs 2>/dev/null; then
	git lfs version
else
	echo "Git LFS is not installed"
	exit
fi

echo "Looks like the prerequisites were installed correctly!"
echo "---"

# Clone and build main repository
go get -v github.com/animenotifier/notify.moe/...
go get -v github.com/stretchr/testify/assert
cd $GOPATH/src/github.com/animenotifier/notify.moe
make all

# Database
git clone https://github.com/animenotifier/database ~/.aero/db/arn

# Configure
make ports

# Add "127.0.0.1 beta.notify.moe" to /etc/hosts
sudo -- sh -c -e "echo '127.0.0.1 beta.notify.moe' >> /etc/hosts"

# Test
make test