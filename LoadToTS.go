package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var rowNumTS = map[string]int{"Key": 0, "Type": 1, "DataStart": 4} // Key:字段key所在行 Type:字段类型所在行 DataStart:数据开始的行 (行数从0开始)

func LoadToTS(excelMap map[string]string) {
	entityHeaderTS := ""
	entityResultTS := ""
	for name, path := range excelMap {
		dataResultMap := map[string]interface{}{}
		entityHeaderTS += "export interface " + name + "   {\n"
		dataMap, entityStrTS := readExcel_TS(path)
		entityResultTS += entityStrTS
		for key, value := range dataMap {
			dataResultMap[key] = value
			entityHeaderTS += ("    " + key + ": { [id: string]: " + key + " };\n")
		}
		byteBuf := bytes.NewBuffer([]byte{})
		encoder := json.NewEncoder(byteBuf)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(dataResultMap)
		if err == nil {
			saveData(byteBuf.String(), outPath+name+".json")
		} else {
			fmt.Println(err.Error())
		}
		entityHeaderTS += "}\n\n"
	}
	saveData(entityHeaderTS+entityResultTS, outPath+"DataEntity.ts")
}

func readExcel_TS(path string) (map[string]map[string]map[string]interface{}, string) {
	file, err := excelize.OpenFile(path)
	if err == nil {
		obj := map[string]map[string]map[string]interface{}{}
		entityStrTS := ""
		sheetMap := file.GetSheetMap()
		for _, sheetName := range sheetMap {
			keyArr := []string{}
			typeArrTS := []string{}
			sheetKey := strings.Title(sheetName)
			obj[sheetKey] = map[string]map[string]interface{}{}
			subObj := obj[sheetKey]
			rows := file.GetRows(sheetName)
			for i, row := range rows {
				for j, cellValue := range row {
					if i == rowNumTS["Key"] {
						keyArr = append(keyArr, cellValue)
					}
					if i == rowNumTS["Type"] {
						typeArrTS = append(typeArrTS, cellValue)
					}
					if i >= rowNumTS["DataStart"] {
						if j == 0 {
							subObj[cellValue] = map[string]interface{}{}
						}
						if typeArrTS[j] == "none" {
							continue
						}
						if typeArrTS[j] != "" {
							subObj[row[0]][keyArr[j]] = getValueByType_TS(cellValue, typeArrTS[j])
						}
					}
				}
			}
			entityStrTS += "export interface " + sheetKey + "  {\n"
			if len(keyArr) == len(typeArrTS) {
				for k, v := range typeArrTS {
					if v == "none" || v == "" {
						continue
					}
					entityStrTS += "    " + keyArr[k] + ": " + typeArrTS[k] + ";\n"
				}
			}
			entityStrTS += "}\n\n"
		}
		return obj, entityStrTS
	} else {
		fmt.Println(err.Error())
	}
	return nil, ""
}

func getValueByType_TS(str string, typeStr string) interface{} {
	str = strings.Trim(str, ";")
	strArr := strings.Split(str, ";")
	var value interface{}
	var arr []interface{}
	if strings.Contains(typeStr, "boolean") {
		if strings.Contains(typeStr, "[]") {
			for _, v := range strArr {
				arr = append(arr, IF(v == "1", true, false))
			}
		} else {
			value = IF(value == "1", true, false)
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
	} else if typeStr == "string[]" {
		for _, v := range strArr {
			arr = append(arr, v)
		}
	} else {
		value = str
	}
	return IF(strings.Contains(typeStr, "[]"), arr, value)
}
