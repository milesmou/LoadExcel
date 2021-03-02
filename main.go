package main

//mac下构建windows命令：GOOS=windows GOARCH=amd64 go build

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var excelMap = map[string]string{}
var outPath string

func walkFunc(path string, info os.FileInfo, err error) error {
	if !info.IsDir() && (strings.Contains(path, "xls") || strings.Contains(path, "xlsx")) {
		base := filepath.Base(filepath.ToSlash(path))
		ext := filepath.Ext(filepath.ToSlash(path))
		name := strings.Title(strings.ReplaceAll(base, ext, ""))
		if !strings.Contains(path, "~") {
			excelMap[name] = filepath.ToSlash(path)
		}
	}
	return nil
}

func main() {
	currentDir := getCurrentDir()
	outPath = currentDir + "/out/"
	os.RemoveAll(outPath)
	filepath.Walk(currentDir, walkFunc)
	LoadToTS(excelMap)
	fmt.Println("Over")
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
