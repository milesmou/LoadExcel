package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var rowNum = map[string]int{"Key": 2, "Type": 1, "DataStart": 4}

var excelMap = map[string][]string{}

func walkFunc(path string, info os.FileInfo, err error) error {
	if !info.IsDir() && (strings.Contains(path, "xls") || strings.Contains(path, "xlsx")) {
		fullDir := filepath.Dir(filepath.ToSlash(path))
		_, dir := filepath.Split(fullDir)
		if dir == "" {
			dir = "default"
		}
		if excelMap[dir] == nil {
			excelMap[dir] = []string{}
		}
		excelMap[dir] = append(excelMap[dir], path)
	}
	return nil
}

func main() {
	entityHeader := ""
	entityResult := ""
	currentDir, _ := os.Getwd()
	filepath.Walk(currentDir, walkFunc)
	for outFileName, pathList := range excelMap {
		dataResultMap := map[string]interface{}{}
		entityHeader += "export interface " + strings.Title(outFileName) + "   {\n"
		for _, path := range pathList {
			dataMap, entityStr := readExcel(path)
			entityResult += entityStr
			for key, value := range dataMap {
				dataResultMap[key] = value
				entityHeader += ("    " + key + ": { [id: number]: " + key + " };\n")
			}
		}
		dataResult, _ := json.Marshal(dataResultMap)
		if dataResult != nil {
			saveData(string(dataResult), currentDir+"\\out\\"+strings.Title(outFileName)+".json")
		}
		entityHeader += "}\n\n"
	}
	saveData(entityHeader+entityResult, currentDir+"\\out\\Entity.ts")
	fmt.Println("按任意键退出")
	fmt.Scanln()
}

func readExcel(path string) (map[string]map[string]map[string]interface{}, string) {
	file, err := excelize.OpenFile(path)
	if err == nil {
		keyArr := []string{}
		typeArr := []string{}
		obj := map[string]map[string]map[string]interface{}{}
		entityStr := ""
		sheetMap := file.GetSheetMap()
		for _, sheetName := range sheetMap {
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
						subObj[row[0]][keyArr[j]] = getValueByType(cellValue, typeArr[j])
					}
				}
			}
			entityStr += "export interface " + sheetKey + "  {\n"
			if len(keyArr) == len(typeArr) {
				for i := 0; i < len(keyArr); i++ {
					entityStr += "    " + keyArr[i] + ": " + typeArr[i] + ";\n"
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
	dir := filepath.Dir(filePath)
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

func IF(condition bool, whenTrue interface{}, whenFalse interface{}) interface{} {
	if condition {
		return whenTrue
	}
	return whenFalse
}
