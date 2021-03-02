package loadts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"

	"LoadExcel/utils"
)

var rowNum = map[string]int{"Key": 0, "Type": 1, "DataStart": 4} // Key:字段key所在行 Type:字段类型所在行 DataStart:数据开始的行 (行数从0开始)
//Excel中的Type支持number、boolean、string以及对应的数组类型

func Load(excelMap map[string]string, outPath string) {
	entityHeader := ""
	entityResult := ""
	for name, path := range excelMap {
		dataResultMap := map[string]interface{}{}
		entityHeader += "export interface " + name + "   {\n"
		dataMap, entityStr := readExcel(path)
		entityResult += entityStr
		for key, value := range dataMap {
			dataResultMap[key] = value
			entityHeader += ("    " + key + ": { [id: string]: " + key + " };\n")
		}
		byteBuf := bytes.NewBuffer([]byte{})
		encoder := json.NewEncoder(byteBuf)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(dataResultMap)
		if err == nil {
			utils.SaveData(byteBuf.String(), outPath+name+".json")
		} else {
			fmt.Println(err.Error())
		}
		entityHeader += "}\n\n"
	}
	utils.SaveData(entityHeader+entityResult, outPath+"DataEntity.ts")
}

func readExcel(path string) (map[string]map[string]map[string]interface{}, string) {
	file, err := excelize.OpenFile(path)
	if err == nil {
		obj := map[string]map[string]map[string]interface{}{}
		entityStr := ""
		sheetMap := file.GetSheetMap()
		for _, sheetName := range sheetMap {
			keyArr := []string{}
			typeArr := []string{}
			sheetKey := strings.Title(sheetName)
			obj[sheetKey] = map[string]map[string]interface{}{}
			subObj := obj[sheetKey]
			rows := file.GetRows(sheetName)
			for i, row := range rows {
				for j, cellValue := range row {
					if i == rowNum["Key"] {
						keyArr = append(keyArr, cellValue)
					}
					if i == rowNum["Type"] {
						typeArr = append(typeArr, cellValue)
					}
					if i >= rowNum["DataStart"] {
						if j == 0 {
							subObj[cellValue] = map[string]interface{}{}
						}
						if typeArr[j] == "none" {
							continue
						}
						if typeArr[j] != "" {
							subObj[row[0]][keyArr[j]] = getValueByType(cellValue, typeArr[j])
						}
					}
				}
			}
			entityStr += "export interface " + sheetKey + "  {\n"
			if len(keyArr) == len(typeArr) {
				for k, v := range typeArr {
					if v == "none" || v == "" {
						continue
					}
					entityStr += "    " + keyArr[k] + ": " + typeArr[k] + ";\n"
				}
			}
			entityStr += "}\n\n"
		}
		return obj, entityStr
	} else {
		fmt.Println(err.Error())
	}
	return nil, ""
}

func getValueByType(str string, typeStr string) interface{} {
	str = strings.Trim(str, ";")
	strArr := strings.Split(str, ";")
	var value interface{}
	var arr []interface{}
	if strings.Contains(typeStr, "boolean") {
		if strings.Contains(typeStr, "[]") {
			for _, v := range strArr {
				arr = append(arr, utils.IF(v == "1", true, false))
			}
		} else {
			value = utils.IF(value == "1", true, false)
		}
	} else if strings.Contains(typeStr, "number") {
		if strings.Contains(typeStr, "[]") {
			for _, v := range strArr {
				v1, _ := strconv.ParseFloat(v, 64)
				arr = append(arr, v1)
			}
		} else {
			v1, _ := strconv.ParseFloat(str, 64)
			value = v1
		}
	} else if strings.Contains(typeStr, "string") {
		if strings.Contains(typeStr, "[]") {
			for _, v := range strArr {
				arr = append(arr, v)
			}
		} else {
			value = str
		}
	} else {
		fmt.Println("不支持的数据类型")
	}
	return utils.IF(strings.Contains(typeStr, "[]"), arr, value)
}
