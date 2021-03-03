package main

//mac下构建windows命令：GOOS=windows GOARCH=amd64 go build

import (
	"os"
	"path/filepath"
	"strings"
)

var excelMap = map[string]string{}
var outPath string

func main() {
	currentDir := getCurrentDir()
	outPath = currentDir + "/out/"
	os.RemoveAll(outPath)
	filepath.Walk(currentDir, walkFunc)
	//LoadTS
	// loadts.Load(excelMap, outPath)
	//LoadCS
	// loadcs.Load(excelMap, outPath)

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
