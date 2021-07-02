#!/usr/bin/python3
import LoadToTS
import LoadToCS

import os
import sys

excelDict: dict = {}
outPath: str = os.getcwd()+os.path.sep+"out"+os.path.sep

# Type:字段类型所在行 Key:字段key所在行 Commit:字段批注所在行 DataStart:数据开始的行 (行数从0开始)
rowNum = {"Type": 1, "Key": 2, "Commit": 3,  "DataStart": 4}


def initExcelDict(currPath):
    for path, listDir, listFile in os.walk(currPath):
        for file in listFile:
            if file.find("~") == -1:  # 排除Excel运行时的临时文件
                fileName: str = os.path.splitext(file)[0]
                fileExt: str = os.path.splitext(file)[-1]
                if fileExt == ".xls" or fileExt == ".xlsx":
                    excelDict[upperFirst(fileName)] = path+os.path.sep+file


def saveData(content: str, fileName: str):
    if not os.path.exists(outPath):
        os.makedirs(outPath)
    with open(outPath+fileName, 'w', encoding="utf-8") as fileIO:
        fileIO.truncate()
        fileIO.write(content)
        fileIO.flush()
        fileIO.close()
        print("数据已保存到->"+outPath+fileName)


def upperFirst(text: str):
    result = ""
    if(text != None and len(text) > 0):
        result = text[0].upper()
        if(len(text) > 1):
            result += text[1:]
    return result


def parseNumber(text: str):
    if(text.find(".") > -1):
        return parseFloat(text)
    else:
        return parseInteger(text)


def parseInteger(text: str):
    value = 0
    try:
        if(text.find(".") > -1):
            text = text[:text.find(".")]
        value = int(text)
    except Exception as e:
        value = 0
    return value


def parseFloat(text: str):
    value = 0.0
    try:
        value = float(text)
    except Exception as e:
        print(e)
    return value


def IF(condition: bool, trueResult, falseResult):
    if condition:
        return trueResult
    else:
        return falseResult


if __name__ == "__main__":
    currPath = os.path.dirname(sys.argv[0])
    print("load path : "+currPath)
    initExcelDict(currPath)
    if(len(excelDict) > 0):
        if(len(sys.argv) >= 2):
            if(sys.argv[1] == "ts"):
                LoadToTS.Load(excelDict)
            elif(sys.argv[1] == "cs"):
                LoadToCS.Load(excelDict)
            else:
                print("参数错误 只支持ts或cs")
        else:
            print("请输入参数ts或cs确定entity类型")
    else:
        print("当前目录及其子目录未找到Excel文件")
