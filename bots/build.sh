#!/bin/sh
MYDIR="$(dirname "$(readlink -f "$0")")"
cd "$MYDIR"
for dir in ./*; do ([ -d "$dir" ] && cd "$dir" && echo "Building bots/$dir" && go build); done