#!/bin/bash

currDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" > /dev/null && pwd)"

cd "$scriptDir/.."

git checkout -- .
git checkout -f master
git pull origin master

/bin/bash scripts/build.sh
/bin/bash scripts/run.sh

cd "$currDir"

exit
