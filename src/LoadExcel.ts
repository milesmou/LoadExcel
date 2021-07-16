import path from "path";
import fs from "fs";
import { LoadToTS } from "./LoadToTS";
import { LoadToCS } from "./LoadToCS";



export class LoadExcel {

    /** Type:字段类型所在行 Key:字段key所在行 Commit:字段批注所在行 DataStart:数据开始的行 (行数从0开始) */
    static rowNum = { "Type": 0, "Key": 1, "Commit": 2, "DataStart": 3 }
    /** 导出文件保存路径 */
    static outPath: string = LoadExcel.getCwd() + path.sep + "out" + path.sep

    static getCwd() {
        let execPath = process.execPath;
        let resolve = path.resolve("./");
        if (execPath.endsWith(path.sep + "node")) {//调试环境
            return resolve;
        } else {//打包后环境
            return path.dirname(execPath);
        }
    }

    static walkDir(dir: string, callback: (file: string) => void) {
        let files = fs.readdirSync(dir)
        files.forEach(v => {
            let filePath = path.join(dir, v)
            if (fs.statSync(filePath).isDirectory()) {
                this.walkDir(path.join(dir, v), callback);
            } else {
                callback(filePath);
            }
        })
    }

    static getXlsxs(): { [name: string]: string } {
        let xlsxs: { [name: string]: string } = {};
        this.walkDir(this.getCwd(), (file: string) => {
            if (file.endsWith(".xlsx") && !path.basename(file).startsWith("~")) {
                let fileName = path.basename(file).replace(path.extname(file), "");
                xlsxs[this.upperFirst(fileName)] = file;
            }
        });
        return xlsxs;
    }

    static upperFirst(str: string): string {
        if (str) {
            return str[0].toUpperCase() + str.slice(1);
        }
        return "";
    }

    static parseNum(str: string) {
        let v = parseFloat(str);
        if (isNaN(v)) return 0;
        return v;
    }

    static saveData(content: string, fileName: string) {
        if (!fs.existsSync(this.outPath)) {
            fs.mkdirSync(this.outPath);
        }
        fs.writeFileSync(path.join(this.outPath, fileName), content);
        console.log("数据已保存到->" + path.join(this.outPath, fileName));
    }


    static loadExcel() {
        console.log("当前路径", this.getCwd());
        let xlsxs = this.getXlsxs();
        if (Object.keys(xlsxs).length > 0) {
            if (process.argv[2] == "ts") {
                LoadToTS.load(xlsxs);
            } else if (process.argv[2] == "cs") {
                LoadToCS.load(xlsxs);
            } else {
                console.log("请输入参数ts或cs确定entity类型");
            }
        } else {
            console.log("当前目录及其子目录未找到Excel文件")
        }
    }
}

LoadExcel.loadExcel();
