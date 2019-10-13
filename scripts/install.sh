#!/bin/bash

currDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" > /dev/null && pwd)"

cd "$scriptDir/.."

mkdir bin

cp -r install/config data/config
ln config/sites.json data/config/sites.json

chmod -R 755 data
chmod -R 755 bin

go get github.com/boltdb/bolt/...

/bin/bash scripts/build.sh

crontab -l > install/crontab.txt

echo "0 3 * * * /bin/bash $scriptDir/update.sh > /dev/null" >> install/crontab.txt
echo "15 * * * * /bin/bash $scriptDir/run.sh > /dev/null" >> install/crontab.txt

crontab install/crontab.txt

rm -f install/crontab.txt

cd "$currDir"

exit
