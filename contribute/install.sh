#!/bin/sh

UBUNTU_VERSION="16.04"
GO_VERSION="1.8.3"
AS_VERSION="3.13.0.1"

GO_FILE="go$GO_VERSION.linux-amd64.tar.gz"
AS_FILE="aerospike-server-community-$AS_VERSION-ubuntu$UBUNTU_VERSION.tgz"
AS_DIR="aerospike-server-community-$AS_VERSION-ubuntu$UBUNTU_VERSION"

if [ ! -d /usr/local/go ]; then
	if [ ! -f "./$GO_FILE" ]; then
		echo "Downloading Go..."
		wget https://storage.googleapis.com/golang/$GOFILE
	fi
	
	echo "Extracting Go..."
	sudo tar -C /usr/local -xzf $GO_FILE

	export PATH=$PATH:/usr/local/go/bin
	echo "Don't forget to add the following to your terminal startup scripts:"
	echo "export PATH=\$PATH:/usr/local/go/bin"
fi

if [ ! -f /usr/bin/asd ]; then
	if [ ! -f ./$AS_FILE ]; then
		echo "Downloading Aerospike..."
		wget http://artifacts.aerospike.com/aerospike-server-community/$AS_VERSION/$AS_FILE
	fi

	if [ ! -d ./$AS_DIR ]; then
		echo "Extracting Aerospike..."
		tar xzf $AS_FILE
	fi

	echo "Installing Aerospike..."
	cd $AS_DIR
	sudo ./asinstall
	cd ..
fi

# if [ ! -d ./notify.moe ]; then
#     echo "Downloading notify.moe source..."
# 	git clone git@github.com:animenotifier/notify.moe.git
# fi

echo "Finished installing notify.moe dependencies."
