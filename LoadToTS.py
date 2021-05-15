from os import fork, name

import LoadExcel
from xlrd.sheet import Sheet
import xlrd
import json

# Key:字段key所在行 Type:字段类型所在行 DataStart:数据开始的行 (行数从0开始)
rowNum = {"Key": 0, "Type": 1, "DataStart": 3}

def Load(excelDict: dict):
    entityHeader: str = ""
    entityResult: str = ""
    for name in excelDict:
        dataResult = {}
        entityHeader += "export interface " + name + "   {\n"
        wbResult = readExcel(excelDict[name])
        entityResult += wbResult["entityStr"]
        for key in wbResult["wbDict"]:
            dataResult[key] = wbResult["wbDict"][key]
            entityHeader += ("    " + key +
                             ": { [id: string]: " + key + " };\n")
        entityHeader += "}\n\n"
        LoadExcel.saveData(json.dumps(
            dataResult, ensure_ascii=False), name+".json")
    LoadExcel.saveData(entityHeader+entityResult, "DataEntity.ts")


def readExcel(path: str):
    with xlrd.open_workbook(path) as workbook:
        wbDict: dict = {}
        entityStr: str = ""
        sheetNames = workbook.sheet_names()
        for sheetName in sheetNames:
            wbDict[str.title(sheetName)] = {}
            sheetDict = wbDict[str.title(sheetName)]
            keyList: list = []
            typeList: list = []
            sheet: Sheet = workbook.sheet_by_name(sheetName)
            for row in range(sheet.nrows):
                idStr = ""
                for col in range(sheet.ncols):
                    cellV = str(sheet.cell_value(row, col)).rstrip("0")
                    cellV = cellV.rstrip(".")
                    if row == rowNum["Key"]:
                        keyList.insert(col, cellV)
                    if row == rowNum["Type"]:
                        typeList.insert(col, cellV)
                    if row >= rowNum["DataStart"]:
                        if col == 0:
                            sheetDict[str(cellV)] = {}
                            idStr = cellV
                        if typeList[col] == "none" or typeList[col] == "":
                            continue
                        sheetDict[idStr][keyList[col]] = getValueByType(
                            cellV, typeList[col])
            entityStr += "export interface " + str.title(sheetName) + "  {\n"
            if len(keyList) == len(typeList):
                for i in range(len(typeList)):
                    v = typeList[i]
                    if v == "none" or v == "":
                        continue
                    entityStr += "    " + \
                        keyList[i] + ": " + typeList[i] + ";\n"
                entityStr += "}\n\n"
        return {"entityStr": entityStr, "wbDict": wbDict}


def getValueByType(cellV: str, typeStr: str):
    cellV = cellV.strip()
    cellV = cellV.strip(";")
    strList = cellV.split(";")
    value = None
    arr: list = []
    if typeStr.find("boolean") > -1:
        if typeStr.find("[]") > -1:
            for v in strList:
                arr.append(LoadExcel.IF(v == 1, True, False))
        else:
            value = LoadExcel.IF(cellV == 1, True, False)
    elif typeStr.find("number") > -1:
        if typeStr.find("[]") > -1:
            for v in strList:
                arr.append(LoadExcel.IF(v.find(".") > -1, float(v), int(v)))
        else:
            value = LoadExcel.IF(cellV.find(".") > -1,
                                 float(cellV), int(cellV))
    elif typeStr.find("string") > -1:
        if typeStr.find("[]") > -1:
            for v in strList:
                arr.append(v)
        else:
            value = cellV
    else:
        print("不支持的数据类型", typeStr)
    return LoadExcel.IF(typeStr.find("[]") > -1, arr, value)