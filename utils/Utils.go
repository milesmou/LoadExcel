package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func SaveData(result string, filePath string) {
	CheckFile(filePath)
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

func CheckFile(filePath string) {
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

func IF(condition bool, whenTrue interface{}, whenFalse interface{}) interface{} {
	if condition {
		return whenTrue
	}
	return whenFalse
}
