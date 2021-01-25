package main

//mac下构建windows命令：GOOS=windows GOARCH=amd64 go build

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var rowNum = map[string]int{"Key": 0, "TypeTS": 1, "TypeCS": 2, "DataStart": 4} // Key:字段key所在行 TypeTS:字段类型所在行 DataStart:数据开始的行 (行数从0开始)

var excelMap = map[string]string{}

func walkFunc(path string, info os.FileInfo, err error) error {
	if !info.IsDir() && (strings.Contains(path, "xls") || strings.Contains(path, "xlsx")) {
		base := filepath.Base(filepath.ToSlash(path))
		ext := filepath.Ext(filepath.ToSlash(path))
		name := strings.Title(strings.ReplaceAll(base, ext, ""))
		excelMap[name] = filepath.ToSlash(path)
	}
	return nil
}

func main() {
	entityHeaderTS := ""
	entityResultTS := ""
	entityHeaderCS := "using System.Collections;\nusing System.Collections.Generic;\n\n"
	entityResultCS := ""
	currentDir := getCurrentDir()
	outPath := currentDir + "/out/"
	os.RemoveAll(outPath)
	filepath.Walk(currentDir, walkFunc)
	fmt.Println("当前路径=" + currentDir)
	for name, path := range excelMap {
		dataResultMap := map[string]interface{}{}
		entityHeaderTS += "export interface " + name + "   {\n"
		entityHeaderCS += "public class " + name + "   \n{\n"
		dataMap, entityStrTS, entityStrCS := readExcel(path)
		entityResultTS += entityStrTS
		entityResultCS += entityStrCS
		for key, value := range dataMap {
			dataResultMap[key] = value
			entityHeaderTS += ("    " + key + ": { [id: number]: " + key + " };\n")
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
		entityHeaderTS += "}\n\n"
		entityHeaderCS += "}\n\n"
	}
	if rowNum["TypeTS"] > 0 {
		saveData(entityHeaderTS+entityResultTS, outPath+"DataEntity.ts")
	}
	if rowNum["TypeCS"] > 0 {
		saveData(entityHeaderCS+entityResultCS, outPath+"DataEntity.cs")
	}
	fmt.Println("Over")
}

func readExcel(path string) (map[string]map[string]map[string]interface{}, string, string) {
	file, err := excelize.OpenFile(path)
	if err == nil {
		obj := map[string]map[string]map[string]interface{}{}
		entityStrTS := ""
		entityStrCS := ""
		sheetMap := file.GetSheetMap()
		for _, sheetName := range sheetMap {
			keyArr := []string{}
			typeArrTS := []string{}
			typeArrCS := []string{}
			sheetKey := strings.Title(sheetName)
			obj[sheetKey] = map[string]map[string]interface{}{}
			subObj := obj[sheetKey]
			rows := file.GetRows(sheetName)
			for i, row := range rows {
				for j, cellValue := range row {
					if i == rowNum["Key"] {
						keyArr = append(keyArr, cellValue)
					}
					if i == rowNum["TypeTS"] {
						typeArrTS = append(typeArrTS, cellValue)
					}
					if i == rowNum["TypeCS"] {
						typeArrCS = append(typeArrCS, cellValue)
					}
					if i >= rowNum["DataStart"] {
						if j == 0 {
							subObj[cellValue] = map[string]interface{}{}
						}
						if typeArrTS[j] == "none" {
							continue
						}
						if typeArrTS[j] != "" {
							subObj[row[0]][keyArr[j]] = getValueByType(cellValue, typeArrTS[j])
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
		return obj, entityStrTS, entityStrCS
	} else {
		fmt.Println(err.Error())
	}
	return nil, "", ""
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
		if os.IsNotExist(err1) {
			e := os.MkdirAll(dir, os.ModeDir)
			if e != nil {
				fmt.Println(e.Error())
			} else {
				os.Chmod(dir, os.ModePerm)
			}
		} else {
			fmt.Println(err1.Error())
		}
	}
	_, err2 := os.Stat(filePath)
	if err2 != nil {
		if os.IsNotExist(err2) {
			_, e := os.Create(filePath)
			if e != nil {
				fmt.Println(e.Error())
			}
		} else {
			fmt.Println(err2.Error())
		}
	}
}

func getCurrentDir() string {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	if strings.Contains(exPath, "go-build") {
		currentDir, _ := os.Getwd()
		return currentDir
	}
	return exPath
}

func IF(condition bool, whenTrue interface{}, whenFalse interface{}) interface{} {
	if condition {
		return whenTrue
	}
	return whenFalse
}
