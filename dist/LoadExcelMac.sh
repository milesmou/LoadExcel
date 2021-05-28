#!/bin/bash
# 使用脚本可以将json和ts,cs文件复制到指定目录

jsonPath=../assets/resources/data/
entityPath=../assets/script/game/

./LoadExcel ts

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
