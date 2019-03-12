#!/bin/bash

echo "Checking prerequisites..."

if hash go 2>/dev/null; then
	if [[ ":$PATH:" == *":$GOPATH/bin:"* || ":$PATH:" == *":$GOPATH/bin/:"* ]]; then
		go version
	else
		echo "Your \$PATH is missing \$GOPATH/bin, you should add this to your ~/.profile:"
		echo "export PATH=\$PATH:/usr/local/go/bin:\$GOPATH/bin"
		exit
	fi
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

# Use sudo here to request permissions for later
sudo echo "---"

# Humanize library is only used in templates
go get -v github.com/dustin/go-humanize

# Clone and build main repository
go get -t -v github.com/animenotifier/notify.moe/...
cd $GOPATH/src/github.com/animenotifier/notify.moe
make all

# Database
git clone --depth=1 https://github.com/animenotifier/database ~/.aero/db/arn

# Configure
make ports

# Add "127.0.0.1 beta.notify.moe" to /etc/hosts
if grep -Fxq "127.0.0.1 beta.notify.moe" /etc/hosts; then
	echo "beta.notify.moe already resolves to localhost"
else
	sudo -- sh -c -e "echo '127.0.0.1 beta.notify.moe' >> /etc/hosts"
	echo "beta.notify.moe now resolves to localhost"
fi

# Test
make test

# Finished
echo "Finished installation."
echo "You can now execute the 'run' command and open https://beta.notify.moe in your browser."
