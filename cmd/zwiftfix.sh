#!/bin/bash

garmin_file=$(ls -1t /media/peter/GARMIN/GARMIN/ACTIVITY/*.FIT | head -1)

echo "GARMIN : $garmin_file"

zwift_file=$(ls -1t ~/Downloads/*.fit | head -1)

echo "ZWIFT : $zwift_file"

cp -v $garmin_file .
cp -v $zwift_file .

echo "go run fitfixer.go ${zwift_file##*/} ${garmin_file##*/}"
go run fitfixer.go ${zwift_file##*/} ${garmin_file##*/}

result=$(ls -1t *.fit | head -1)

echo "Ready : $result"
