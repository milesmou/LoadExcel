import xlsx from "xlsx"
import { LoadExcel } from "./LoadExcel";

export class LoadToTS {

    static load(xlsxs: { [name: string]: string }) {
        let entityHeader = "";
        let entityResult = "";
        for (const name in xlsxs) {
            let dataResult: { [name: string]: object } = {};
            entityHeader += "export interface " + name + "   {\n";
            let wbResult = this.readExcel(xlsxs[name]);
            entityResult += wbResult["entityStr"];
            for (const key in wbResult["wbDict"]) {
                dataResult[key] = wbResult["wbDict"][key];
                entityHeader += ("    " + key + ": { [id: string]: " + key + " };\n");
            }
            entityHeader += "}\n\n";
            LoadExcel.saveData(JSON.stringify(dataResult), name + ".json");
        }
        LoadExcel.saveData(entityHeader + entityResult, "DataEntity.ts");
    }

    static readExcel(filePath: string): { entityStr: string, wbDict: { [name: string]: { [key: string]: any } } } {
        let workbook = xlsx.readFile(filePath, { type: "array" });
        let entityStr: string = "";
        let wbDict: { [name: string]: { [key: string]: any } } = {};
        for (const sheetName of workbook.SheetNames) {
            if (sheetName.startsWith("~")) continue;
            let sheet = workbook.Sheets[sheetName];
            let sName = LoadExcel.upperFirst(sheetName);
            wbDict[sName] = {};
            let sheetDict = wbDict[sName];
            let typeList: string[] = [];
            let keyList: string[] = [];
            let commitList: string[] = [];
            let rowsData: string[][] = xlsx.utils.sheet_to_json(sheet, { header: 1, defval: "" });
            console.log("load sheet " + sName + " start");
            for (let row = 0; row < rowsData.length; row++) {
                let id = "";
                const colsData = rowsData[row];
                for (let col = 0; col < colsData.length; col++) {
                    const cellV = colsData[col].toString();
                    if (row == LoadExcel.rowNum.Type) {
                        typeList[col] = cellV.replace(/"/g, "");
                    }
                    if (row == LoadExcel.rowNum.Key) {
                        keyList[col] = cellV.replace(/"/g, "");
                    }
                    if (row == LoadExcel.rowNum.Commit) {
                        commitList[col] = cellV;
                    }
                    if (row >= LoadExcel.rowNum.DataStart) {
                        if (col == 0 && cellV) {
                            sheetDict[cellV] = {};
                            id = cellV.replace(/"/g, "");
                        }
                        let type = typeList[col];
                        let key = keyList[col];
                        if (!id || !type || type == "none" || !key) continue;
                        sheetDict[id][key] = this.getValueByType(cellV, type);
                    }
                }
            }
            entityStr += "export class " + sName + "  {\n";
            for (let i = 0; i < typeList.length; i++) {
                let type = typeList[i];
                let key = keyList[i];
                if (!type || type == "none" || !key) continue;
                entityStr += "    /** " + commitList[i] + " */\n";
                entityStr += "    " + key + ": " + type + ";\n";
            }
            entityStr += "}\n\n";
            console.log("load sheet " + sName + " end");
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