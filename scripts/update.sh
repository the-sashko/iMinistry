#!/bin/bash

currDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" > /dev/null && pwd)"

cd "$scriptDir/.."

git checkout -- .
git checkout -f master
git pull origin master

rm -f data/config/sites.json
cp config/sites.json data/config/sites.json

/bin/bash scripts/build.sh

cd "$currDir"

exit
