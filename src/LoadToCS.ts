import excel from "exceljs";
import { LoadExcel } from "./LoadExcel";

export class LoadToCS {
    static async load(xlsxs: { [name: string]: string }) {
        let entityHeader = "using System.Collections;\nusing System.Collections.Generic;\n\n";
        let entityResult = "";
        for (const name in xlsxs) {
            let dataResult: { [name: string]: object } = {}
            entityHeader += "public class " + name + "   \n{\n"
            let wbResult = await this.readExcel(xlsxs[name])
            entityResult += wbResult["entityStr"]
            for (const key in wbResult["wbDict"]) {
                dataResult[key] = wbResult["wbDict"][key]
                entityHeader += ("    public Dictionary<string," + key + "> " + key + ";\n")
            }
            entityHeader += "}\n\n"
            LoadExcel.saveData(JSON.stringify(dataResult), name + ".json")
        }
        LoadExcel.saveData(entityHeader + entityResult, "DataEntity.cs")
    }

    static async readExcel(filePath: string): Promise<{ entityStr: string, wbDict: { [name: string]: { [key: string]: any } } }> {
        let workbook = new excel.Workbook();
        await workbook.xlsx.readFile(filePath)
        let sheets = workbook.worksheets;
        let wbDict: { [name: string]: { [key: string]: any } } = {}
        let entityStr: string = ""
        for (const sheet of sheets) {
            console.log("load sheet " + sheet.name + " start")
            let sName = LoadExcel.upperFirst(sheet.name)
            wbDict[sName] = {}
            let sheetDict = wbDict[sName]
            let typeList: string[] = []
            let keyList: string[] = []
            let commitList: string[] = []
            for (let row = 1; row < sheet.rowCount; row++) {
                let id = ""
                for (let col = 1; col < sheet.columnCount; col++) {
                    let value = sheet.getCell(row, col).value;
                    let cellV = value ? value.toString() : "";
                    if (row == LoadExcel.rowNum.Type) {
                        typeList[col] = cellV;
                    }
                    if (row == LoadExcel.rowNum.Key) {
                        keyList[col] = cellV;
                    }
                    if (row == LoadExcel.rowNum.Commit) {
                        commitList[col] = cellV;
                    }
                    if (row >= LoadExcel.rowNum.DataStart) {
                        if (col == 1 && cellV) {
                            sheetDict[cellV] = {}
                            id = cellV;
                        }
                        if (!id || !typeList[col] || typeList[col] == "none" || !keyList[col]) continue;
                        sheetDict[id][keyList[col]] = this.getValueByType(cellV, typeList[col])
                    }
                }
            }
            entityStr += "public class " + sName + "  \n{\n"
            if (keyList.length == typeList.length) {
                for (let i = 0; i < typeList.length; i++) {
                    let type = typeList[i]
                    if (!type || type == "none") continue;
                    entityStr += "    /// <summary>\n"
                    entityStr += "    /// " + commitList[i] + "\n"
                    entityStr += "    /// <summary>\n"
                    entityStr += "    public " + typeList[i] + " " + keyList[i] + ";\n"
                }
            }
            entityStr += "}\n\n"
            console.log("load sheet " + sheet.name + " end")
        }
        return { "entityStr": entityStr, "wbDict": wbDict };
    }

    static getValueByType(cellV: string, type: string): any {
        cellV = cellV.trim()
        let result: any;
        if (type.includes("bool")) {
            if (type == "bool[][]") {
                result = cellV.split(";").filter(v => v).map(v => v.split(",").filter(v => v).map(v => v == "1" ? true : false));
            } else if (type == "bool[]") {
                result = cellV.split(";").filter(v => v).map(v => v == "1" ? true : false);
            } else {
                result = cellV == "1" ? true : false;
            }
        } else if (type.includes("int") || type.includes("float")) {
            if (type == "int[][]" || type == "float[][]") {
                result = cellV.split(";").filter(v => v).map(v => v.split(",").filter(v => v).map(v => LoadExcel.parseNum(v)));
            } else if (type == "int[]" || type == "float[]") {
                result = cellV.split(";").filter(v => v).map(v => LoadExcel.parseNum(v));
            } else {
                result = LoadExcel.parseNum(cellV)
            }
        } else if (type.includes("string")) {
            if (type == "string[][]") {
                result = cellV.split(";").filter(v => v).map(v => v.split(",").filter(v => v));
            } else if (type == "string[]") {
                result = cellV.split(";").filter(v => v);
            } else {
                result = cellV
            }
        } else {
            console.log("不支持的数据类型", type)
        }
        return result;
    }
}