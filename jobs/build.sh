#!/bin/sh
MYDIR="$(dirname "$(realpath "$0")")"
cd "$MYDIR"
for dir in ./*; do ([ -d "$dir" ] && cd "$dir" && echo "Building jobs/$dir" && go build); done