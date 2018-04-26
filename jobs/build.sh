#!/bin/sh
MYDIR="$(dirname "$(readlink -f "$0")")"
cd "$MYDIR"
go build
for dir in ./*; do ([ -d "$dir" ] && cd "$dir" && echo "Building jobs/$dir" && go build); done