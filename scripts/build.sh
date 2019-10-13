#!/bin/bash

currDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" > /dev/null && pwd)"

cd "$scriptDir/.."

go build -o bin/run src/*

chmod -X bin/run
chmod 755 bin/run

/bin/bash scripts/run.sh

cd "$currDir"

exit
