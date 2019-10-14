#!/bin/bash

currDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" > /dev/null && pwd)"

cd "$scriptDir/../docker"

sudo docker build -t i_ministry_bot .

cd "$currDir"

mkdir data

sudo docker run -d -p 80:80 --name i_ministry_bot -v "$(pwd)/data":/storage/tgBot/data i_ministry_bot

cp "$scriptDir/cron.sh" cron.sh

exit
