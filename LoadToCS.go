package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var rowNumCS = map[string]int{"Key": 0, "Type": 1, "DataStart": 4} // Key:字段key所在行 Type:字段类型所在行 DataStart:数据开始的行 (行数从0开始)

func LoadToCS(excelMap map[string]string) {
	entityHeaderCS := "using System.Collections;\nusing System.Collections.Generic;\n\n"
	entityResultCS := ""
	for name, path := range excelMap {
		dataResultMap := map[string]interface{}{}
		entityHeaderCS += "public class " + name + "   \n{\n"
		dataMap, entityStrCS := readExcel_CS(path)
		entityResultCS += entityStrCS
		for key, value := range dataMap {
			dataResultMap[key] = value
			entityHeaderCS += ("    public Dictionary<string," + key + "> " + key + ";\n")
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
		entityHeaderCS += "}\n\n"
	}
	saveData(entityHeaderCS+entityResultCS, outPath+"DataEntity.cs")
}

func readExcel_CS(path string) (map[string]map[string]map[string]interface{}, string) {
	file, err := excelize.OpenFile(path)
	if err == nil {
		obj := map[string]map[string]map[string]interface{}{}
		entityStrCS := ""
		sheetMap := file.GetSheetMap()
		for _, sheetName := range sheetMap {
			keyArr := []string{}
			typeArrCS := []string{}
			sheetKey := strings.Title(sheetName)
			obj[sheetKey] = map[string]map[string]interface{}{}
			subObj := obj[sheetKey]
			rows := file.GetRows(sheetName)
			for i, row := range rows {
				for j, cellValue := range row {
					if i == rowNumCS["Key"] {
						keyArr = append(keyArr, cellValue)
					}
					if i == rowNumCS["Type"] {
						typeArrCS = append(typeArrCS, cellValue)
					}
					if i >= rowNumCS["DataStart"] {
						if j == 0 {
							subObj[cellValue] = map[string]interface{}{}
						}
						if typeArrCS[j] == "none" {
							continue
						}
						if typeArrCS[j] != "" {
							subObj[row[0]][keyArr[j]] = getValueByType_CS(cellValue, typeArrCS[j])
						}
					}
				}
			}
			entityStrCS += "public class " + sheetKey + "  \n{\n"
			if len(keyArr) == len(typeArrCS) {
				for k, v := range typeArrCS {
					if v == "none" || v == "" {
						continue
					}
					entityStrCS += "    public " + typeArrCS[k] + " " + keyArr[k] + ";\n"
				}
			}
			entityStrCS += "}\n\n"
		}
		return obj, entityStrCS
	} else {
		fmt.Println(err.Error())
	}
	return nil, ""
}

func getValueByType_CS(str string, typeStr string) interface{} {
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
