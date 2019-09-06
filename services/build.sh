#!/bin/sh
ROOT=$PWD
INSTALLPATH="/etc/systemd/system/"
GOBIN="$(go env GOPATH)/bin"
cd $(dirname $0)

for service in *.service; do
	[ -f "$service" ] &&
	echo "Installing services/$service" &&
	sudo cp "$service" "$INSTALLPATH/$service" &&
	sudo sed -i "s|MAKEFILE_USER|$USER|g" "$INSTALLPATH/$service" &&
	sudo sed -i "s|MAKEFILE_GOBIN|$GOBIN|g" "$INSTALLPATH/$service" &&
	sudo sed -i "s|MAKEFILE_PWD|$ROOT|g" "$INSTALLPATH/$service"
done

sudo systemctl daemon-reload &&
echo -e "\nYou can now start the service using:\n\nsudo systemctl start animenotifier.service"