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
	dirs := make([]string, 0)
	listDir("../..", dirs)
}

func listDir(dirpath string, dirs []string) {
	fileInfos, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, subFile := range fileInfos {
		if subFile.IsDir() && !IsFileHidden(subFile) {
			fileDir := dirpath + "/" + subFile.Name()
			dirs = append(dirs, fileDir)

			fmt.Println("目录:", fileDir)
			listDir(fileDir, dirs)
		} else {
			fmt.Println("文件:", subFile.Name())
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
