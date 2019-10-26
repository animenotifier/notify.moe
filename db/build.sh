#!/bin/sh
TYPES=`curl -sS 'https://notify.moe/api/types' | gojq '.[]' | cut -d '"' -f 2`
SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"
cd "$SCRIPTPATH"
mkdir -p arn

for TYPENAME in $TYPES; do
	echo "Downloading db/arn/$TYPENAME.dat"
	URL=https://notify.moe/api/types/$TYPENAME/download
	curl -sS $URL -o arn/$TYPENAME.dat
done
