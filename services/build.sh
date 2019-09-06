#!/bin/sh
INSTALLPATH="/etc/systemd/system/"
cd $(dirname $0)

for service in *; do
	[ -f "$service" ] &&
	echo "Installing services/$service" &&
	sudo cp "services/$service" "$INSTALLPATH/$service" &&
	sudo sed -i "s|MAKEFILE_USER|$USER|g" "$INSTALLPATH/$service" &&
	sudo sed -i "s|MAKEFILE_PWD|$PWD|g" "$INSTALLPATH/$service"
done

sudo systemctl daemon-reload
echo -e "\nYou can now start the service using:\n\nsudo systemctl start animenotifier.service"