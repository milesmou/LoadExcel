import LoadExcel
from xlrd.sheet import Sheet
import xlrd
import json

# Key:字段key所在行 Type:字段类型所在行 DataStart:数据开始的行 (行数从0开始)
rowNum = {"Key": 2, "Type": 1, "DataStart": 4}

def Load(excelDict: dict):
    entityHeader: str = "using System.Collections;\nusing System.Collections.Generic;\n\n"
    entityResult: str = ""
    for name in excelDict:
        dataResult = {}
        entityHeader += "public class " + name + "   \n{\n"
        wbResult = readExcel(excelDict[name])
        entityResult += wbResult["entityStr"]
        for key in wbResult["wbDict"]:
            dataResult[key] = wbResult["wbDict"][key]
            entityHeader += ("    public Dictionary<string," + key + "> " + key + ";\n")
        entityHeader += "}\n\n"
        LoadExcel.saveData(json.dumps(dataResult, ensure_ascii=False), name+".json")
    LoadExcel.saveData(entityHeader+entityResult, "DataEntity.cs")


def readExcel(path: str):
    with xlrd.open_workbook(path) as workbook:
        wbDict: dict = {}
        entityStr: str = ""
        sheetNames = workbook.sheet_names()
        for sheetName in sheetNames:
            print("load sheet "+sheetName+" start")
            sName = LoadExcel.upperFirst(sheetName)
            wbDict[sName] = {}
            sheetDict = wbDict[sName]
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
            entityStr += "public class " + sName + "  \n{\n"
            if len(keyList) == len(typeList):
                for i in range(len(typeList)):
                    v = typeList[i]
                    if v == "none" or v == "":
                        continue
                    entityStr += "    public " + typeList[i] + " " + keyList[i] + ";\n"
                entityStr += "}\n\n"
            print("load sheet "+sheetName+" end")
        return {"entityStr": entityStr, "wbDict": wbDict}


def getValueByType(cellV: str, typeStr: str):
    cellV = cellV.strip()
    cellV = cellV.strip(";")
    strList = cellV.split(";")
    value = None
    arr: list = []
    if typeStr.find("bool") > -1:
        if typeStr.find("[]") > -1:
            for v in strList:
                arr.append(LoadExcel.IF(v == 1, True, False))
        else:
            value = LoadExcel.IF(cellV == 1, True, False)
    elif typeStr.find("int") > -1 or typeStr.find("float") > -1:
        if typeStr.find("[]") > -1:
            for v in strList:
                try:
                    arr.append(LoadExcel.IF(typeStr.find("float") > -1, float(v), int(v)))
                except Exception:
                    arr.append(0)
        else:
            try:
                value = LoadExcel.IF(typeStr.find("float") > -1, float(cellV), int(cellV))
            except Exception:
                value = 0

    elif typeStr.find("string") > -1:
        if typeStr.find("[]") > -1:
            for v in strList:
                arr.append(v)
        else:
            value = cellV
    else:
        print("不支持的数据类型", typeStr)
    return LoadExcel.IF(typeStr.find("[]") > -1, arr, value)
