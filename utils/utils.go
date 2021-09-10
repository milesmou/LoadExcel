package utils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

const outPath = "./out/"

var IsDebug bool

func Println(str string) {
	if IsDebug {
		fmt.Println(str)
	}
}

func IF(condition bool, trueResult interface{}, falseResult interface{}) interface{} {
	if condition {
		return trueResult
	} else {
		return falseResult
	}
}

func ClearOut() {
	if _, err := os.Stat(outPath); err == nil {
		if err := os.RemoveAll(outPath); err != nil {
			fmt.Println("清空上次导出文件失败")
		}
	}
	if err := os.MkdirAll(outPath, os.ModePerm); err != nil {
		fmt.Println("创建out目录失败")
	}
}

func SaveDataWithMap(value map[string]interface{}, fileName string) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := jsoniter.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(value)
	SaveDataWithString(bf.String(), fileName)
}

func SaveDataWithString(value string, fileName string) {
	file, err := os.OpenFile(filepath.Join(outPath, fileName), os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err == nil {
		defer file.Close()
		if _, err := file.Write([]byte(value)); err == nil {
			fmt.Println("数据已保存到->" + filepath.Join(outPath, fileName))
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("数据保存失败")
	}
}

func ParseNum(str string) interface{} {
	if strings.Contains(str, ".") {
		return ParseFloat(str)
	} else {
		return ParseInt(str)
	}
}

func ParseInt(str string) int {
	v, e := strconv.Atoi(str)
	if e != nil {
		v = 0
	}
	return v
}

func ParseFloat(str string) float64 {
	v, e := strconv.ParseFloat(str, 64)
	if e != nil {
		v = 0
	}
	return v
}
