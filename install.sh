#!/bin/bash

# Clone and build main repository
go get -v github.com/animenotifier/notify.moe/...
cd $GOPATH/src/github.com/animenotifier/notify.moe
make all

# Database
git clone https://github.com/animenotifier/database ~/.aero/db/arn

# Configure
make ports
sudo -- sh -c -e "echo '127.0.0.1 beta.notify.moe' >> /etc/hosts"
