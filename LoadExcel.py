#!/usr/bin/python3

import LoadToTS
import LoadToCS

import os

excelDict: dict = {}
outPath: str = os.getcwd()+os.path.sep+"Out"+os.path.sep


def initExcelDict():
    for path, listDir, listFile in os.walk(os.getcwd()):
        for file in listFile:
            if file.find("~") == -1:  # 排除Excel运行时的临时文件
                fileName: str = os.path.splitext(file)[0]
                fileExt: str = os.path.splitext(file)[-1]
                if fileExt == ".xls" or fileExt == ".xlsx":
                    excelDict[fileName.title()] = path+os.path.sep+file


def saveData(content: str, fileName: str):
    if not os.path.exists(outPath):
        os.makedirs(outPath)
    with open(outPath+fileName, 'w') as fileIO:
        fileIO.truncate()
        fileIO.write(content)
        fileIO.flush()
        fileIO.close()
        print("数据已保存到->"+outPath+fileName)


def IF(condition: bool, trueResult, falseResult):
    if condition:
        return trueResult
    else:
        return falseResult


if __name__ == "__main__":
    initExcelDict()
    # LoadToTS.Load(excelDict)
    LoadToCS.Load(excelDict)
