#!/bin/bash

currDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" > /dev/null && pwd)"

cd "$scriptDir/.."

mkdir bin
mkdir -p data/config

cp -r install/config/credentials.json data/config/credentials.json
cp config/sites.json data/config/sites.json

chmod -R 755 data
chmod -R 755 bin

go get github.com/boltdb/bolt/...

/bin/bash scripts/build.sh

cd "$currDir"

exit
