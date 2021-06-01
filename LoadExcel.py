#!/usr/bin/python3
import LoadToTS
import LoadToCS

import os
import sys

excelDict: dict = {}
outPath: str = os.getcwd()+os.path.sep+"out"+os.path.sep


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
