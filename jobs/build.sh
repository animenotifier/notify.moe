#!/bin/sh
SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"
cd "$SCRIPTPATH"
go build

for dir in *; do
	[ -d "$SCRIPTPATH/$dir" ] &&
	cd "$SCRIPTPATH/$dir" &&
	echo "Building jobs/$dir" &&
	go build
done