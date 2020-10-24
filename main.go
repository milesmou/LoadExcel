package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var outPath = map[string]string{"Data": "./out/Data.json", "Entity": "./out/Entity.ts"}

var rowNum = map[string]int{"Key": 2, "Type": 1, "DataStart": 4}

var excelList []string

func walkFunc(path string, info os.FileInfo, err error) error {
	if !info.IsDir() && (strings.Contains(path, "xls") || strings.Contains(path, "xlsx")) {
		excelList = append(excelList, path)
	}
	return nil
}

func main() {
	dataResultMap := map[string]interface{}{}
	entityHeader := "export interface IJsonData   {\n"
	entityResult := ""
	currentDir, _ := os.Getwd()
	filepath.Walk(currentDir, walkFunc)
	for i := 0; i < len(excelList); i++ {
		filePath := excelList[i]
		index := strings.LastIndex(filePath, "\\")
		if index == -1 {
			index = strings.LastIndex(filePath, "/")
		}
		fileName := filePath[index+1 : len(filePath)-len(path.Ext(filePath))]
		dataMap, entityStr := readExcel(filePath, fileName)
		dataResultMap[strings.Title(fileName)] = dataMap
		entityResult += entityStr
		entityHeader += ("    " + strings.Title(fileName) + ": { [id: number]: " + strings.Title(fileName) + " };\n")
	}
	entityHeader += "}\n\n"
	dataResult, _ := json.Marshal(dataResultMap)
	if dataResult != nil {
		saveData(string(dataResult), outPath["Data"])
	}
	saveData(entityHeader+entityResult, outPath["Entity"])
}

func readExcel(path string, name string) (map[string]map[string]interface{}, string) {
	file, err := excelize.OpenFile(path)
	if err == nil {
		keyArr := []string{}
		typeArr := []string{}
		obj := map[string]map[string]interface{}{}
		sheetMap := file.GetSheetMap()
		for _, sheetName := range sheetMap {
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
							obj[cellValue] = map[string]interface{}{}
						}
						obj[row[0]][keyArr[j]] = getValueByType(cellValue, typeArr[j])
					}
				}
			}
		}
		entityStr := "export interface " + strings.Title(name) + "  {\n"
		if len(keyArr) == len(typeArr) {
			for i := 0; i < len(keyArr); i++ {
				entityStr += "    " + keyArr[i] + ": " + typeArr[i] + ";\n"
			}
		}
		entityStr += "}\n\n"
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
				if v == "0" {
					arr = append(arr, false)
				} else {
					arr = append(arr, true)
				}
			}
		} else {
			if str == "0" {
				value = false
			} else {
				value = true
			}
		}
	} else if strings.Contains(typeStr, "number") {
		if strings.Contains(typeStr, "[]") {
			for _, v := range strArr {
				if strings.Contains(v, ".") {
					v2, _ := strconv.ParseFloat(v, 64)
					arr = append(arr, v2)
				} else {
					v2, _ := strconv.ParseInt(v, 10, 64)
					arr = append(arr, v2)
				}
			}
		} else {
			if strings.Contains(str, ".") {
				v, _ := strconv.ParseFloat(str, 64)
				value = v
			} else {
				v, _ := strconv.ParseInt(str, 10, 64)
				value = v
			}
		}
	} else {
		value = str
	}
	if strings.Contains(typeStr, "[]") {
		return arr
	} else {
		return value
	}
}

func saveData(result string, filePath string) {
	checkFile(filePath)
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0o666)
	if err == nil {
		file.Truncate(0)
		file.WriteString(result)
		file.Close()
		fmt.Println("数据已保存到->" + filePath)
	} else {
		fmt.Println(err.Error())
	}
}

func checkFile(filePath string) {
	dir := "./" + path.Dir(filePath)
	_, err1 := os.Stat(dir)
	if err1 != nil {
		e := os.Mkdir(dir, os.ModeDir)
		if e != nil {
			fmt.Println(e.Error())
		}
	}
	_, err2 := os.Stat(filePath)
	if err2 != nil {
		_, e := os.Create(filePath)
		if e != nil {
			fmt.Println(e.Error())
		}
	}
}
