#!/bin/bash

currDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" > /dev/null && pwd)"

cd "$scriptDir/"

./update.sh

cd "$scriptDir/../bin"

./run

cd "$currDir"

exit
