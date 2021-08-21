package main

import (
	"LoadExcel/loadToCS"
	"LoadExcel/loadToTS"
	"LoadExcel/utils"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	xlsxs := getXlsxs()
	utils.ClearOut()
	if len(xlsxs) == 0 {
		fmt.Println("当前目录及其子目录未找到Excel文件")
		return
	}
	if len(os.Args) >= 3 {
		utils.IsDebug = true
	}
	if len(os.Args) >= 2 {
		v := os.Args[1]
		if v == "ts" {
			loadToTS.Load(xlsxs)
		} else if v == "cs" {
			loadToCS.Load(xlsxs)
		} else {
			fmt.Println("请输入参数ts或cs确定entity类型")
		}
	} else {
		fmt.Println("请输入参数ts或cs确定entity类型")
	}
}

func getXlsxs() map[string]string {
	xlsxs := map[string]string{}
	if curPath, err := os.Getwd(); err == nil {
		fmt.Printf("当前路径：%s\n", curPath)
		filepath.Walk(curPath, func(path string, info fs.FileInfo, err error) error {
			if err == nil {
				ext := filepath.Ext(path)
				name := strings.Title(strings.Replace(filepath.Base(path), ext, "", -1))
				if ext == ".xlsx" || ext == ".xls" {
					if !strings.HasPrefix(name, "~") {
						xlsxs[name] = path
					}
				}
				return nil
			} else {
				return err
			}
		})
	}
	return xlsxs
}
