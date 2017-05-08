package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {
	fileInfos, err := ioutil.ReadDir("..")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() && fileInfo.Name() != "clean" && !IsFileHidden(fileInfo) {
			fmt.Println(fileInfo.Name())
		}
	}
}

func IsFileHidden(file os.FileInfo) bool {
	fa := reflect.ValueOf(file.Sys()).Elem().FieldByName("FileAttributes").Uint()
	bytefa := []byte(strconv.FormatUint(fa, 2))
	if bytefa[len(bytefa)-2] == '1' {
		//fmt.Println("隐藏目录:", file.Name())
		return true
	}
	return false
}
