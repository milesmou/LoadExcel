#!/bin/bash

jsonPath=../assets/resources/data/
entityPath=../assets/script/game/

./LoadExcel

files=$(ls ./out)
for filename in $files; do
    ext=${filename#*.}
    if [ $ext == "json" ]; then
        cp -R ./out/$filename $jsonPath
    fi
    if [ $ext == "cs" ] || [ $ext == "ts" ]; then
        cp -R ./out/$filename $entityPath
    fi
done
