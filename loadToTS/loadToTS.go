package loadToTS

import (
	"LoadExcel/utils"
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const (
	TypeRow   = 0
	KeyRow    = 1
	CommitRow = 2
	DataStart = 3
)

func Load(xlsxs map[string]string) {
	var entityHeader = ""
	var entityResult = ""
	for name, path := range xlsxs {
		entityHeader += "export interface " + name + "   {\n"
		entity, obj := readExcel(path)
		entityResult += entity
		for k := range obj {
			entityHeader += "    " + k + ": { [id: string]: " + k + " };\n"
		}
		utils.SaveDataWithMap(obj, name+".json")
		entityHeader += "}\n\n"
	}
	utils.SaveDataWithString(entityHeader+entityResult, "DataEntity.ts")
}

func readExcel(path string) (string, map[string]interface{}) {
	var entityStr = ""
	var wbMap = map[string]interface{}{}
	if file, err := excelize.OpenFile(path); err == nil {
		for i := 1; i <= file.SheetCount; i++ {
			sheetName := file.GetSheetName(i)
			if strings.HasPrefix(sheetName, "~") {
				continue
			}
			utils.Println("load sheet " + sheetName + " start")
			sName := strings.Title(sheetName)
			typeList := []string{}
			keyList := []string{}
			commitList := []string{}
			sheetMap := map[string]map[string]interface{}{}
			for row, cols := range file.GetRows(sheetName) {
				id := ""
				for col, cell := range cols {
					if row == TypeRow {
						typeList = append(typeList, cell)
					}
					if row == KeyRow {
						keyList = append(keyList, cell)
					}
					if row == CommitRow {
						commitList = append(commitList, cell)
					}
					if row >= DataStart {
						if col == 0 && cell != "" {
							id = cell
							sheetMap[id] = map[string]interface{}{}
						}
						t := typeList[col]
						k := keyList[col]
						if t == "" || t == "none" {
							continue
						}
						if sheetMap[id] != nil {
							sheetMap[id][k] = getValueByType(cell, t)
						}
					}
				}
			}
			wbMap[sName] = sheetMap
			entityStr += "export class " + sName + "  {\n"
			for i := 0; i < len(typeList); i++ {
				t := typeList[i]
				k := keyList[i]
				if t == "" || t == "none" || k == "" {
					continue
				}
				entityStr += "    /** " + commitList[i] + " */\n"
				entityStr += "    " + k + ": " + t + ";\n"
			}
			entityStr += "}\n\n"
			utils.Println("load sheet " + sheetName + " end")
		}
	}
	return entityStr, wbMap
}

func getValueByType(v string, t string) interface{} {
	var result interface{}
	v = strings.Trim(v, " ")
	v = strings.Trim(v, ";")
	if strings.HasPrefix(t, "boolean") {
		if t == "boolean[][]" {
			result = utils.ParseArrArr(v, ";", ",", func(v1 string) interface{} {
				return utils.IF(v1 == "" || v1 == "0", false, true)
			})
		} else if t == "boolean[]" {
			result = utils.ParseArr(v, ";", func(v1 string) interface{} {
				return utils.IF(v1 == "" || v1 == "0", false, true)
			})
		} else {
			result = utils.IF(v == "" || v == "0", false, true)
		}
	} else if strings.HasPrefix(t, "number") {
		if t == "number[][]" {
			result = utils.ParseArrArr(v, ";", ",", utils.ParseNum)
		} else if t == "number[]" {
			result = utils.ParseArr(v, ";", utils.ParseNum)
		} else {
			result = utils.ParseNum(v)
		}
	} else if strings.HasPrefix(t, "string") {
		if t == "string[][]" {
			result = utils.ParseArrArr(v, ";", ",", func(v1 string) interface{} { return v1 })
		} else if t == "string[]" {
			result = utils.ParseArr(v, ";", func(v1 string) interface{} { return v1 })
		} else {
			result = v
		}
	} else {
		fmt.Println("未知的数据类型", t)
	}
	return result
}
