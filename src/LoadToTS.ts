import excel from "exceljs";
import { LoadExcel } from "./LoadExcel";

export class LoadToTS {

    static async load(xlsxs: { [name: string]: string }) {
        let entityHeader = "";
        let entityResult = "";
        for (const name in xlsxs) {
            let dataResult: { [name: string]: object } = {}
            entityHeader += "export interface " + name + "   {\n"
            let wbResult = await this.readExcel(xlsxs[name])
            entityResult += wbResult["entityStr"]
            for (const key in wbResult["wbDict"]) {
                dataResult[key] = wbResult["wbDict"][key]
                entityHeader += ("    " + key + ": { [id: string]: " + key + " };\n")
            }
            entityHeader += "}\n\n"
            LoadExcel.saveData(JSON.stringify(dataResult), name + ".json")
        }
        LoadExcel.saveData(entityHeader + entityResult, "DataEntity.ts")
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
            for (let row = 1; row <= sheet.rowCount; row++) {
                let id = ""
                for (let col = 1; col <= sheet.columnCount; col++) {
                    let cellV = sheet.getCell(row, col).toString();
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
            entityStr += "export interface " + sName + "  {\n"
            if (keyList.length == typeList.length) {
                for (let i = 0; i < typeList.length; i++) {
                    let type = typeList[i]
                    if (!type || type == "none") continue;
                    entityStr += "    /** " + commitList[i] + " */\n"
                    entityStr += "    " + keyList[i] + ": " + typeList[i] + ";\n"
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
        if (type.includes("boolean")) {
            if (type == "boolean[][]") {
                result = cellV.split(";").filter(v => v).map(v => v.split(",").filter(v => v).map(v => v == "1" ? true : false));
            } else if (type == "boolean[]") {
                result = cellV.split(";").filter(v => v).map(v => v == "1" ? true : false);
            } else {
                result = cellV == "1" ? true : false;
            }
        } else if (type.includes("number")) {
            if (type == "number[][]") {
                result = cellV.split(";").filter(v => v).map(v => v.split(",").filter(v => v).map(v => LoadExcel.parseNum(v)));
            } else if (type == "number[]") {
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