#!/bin/sh
SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"
cd "$SCRIPTPATH"

for dir in *; do
	[ -d "$SCRIPTPATH/$dir" ] &&
	cd "$SCRIPTPATH/$dir" &&
	echo "Building bots/$dir" &&
	go build
done