package main

import (
	"flag"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"os"
	"path/filepath"
)

var confPath *string

func init() {
	confPath = flag.String("c", "../resources", "conf dir")
	flag.Parse()
}

func main() {
	workPath, _ := os.Getwd()
	newPath := filepath.Join(workPath, *confPath)
	fileInfos, _ := ioutil.ReadDir(newPath)
	fmt.Println("out1: ", newPath)
	fmt.Println("out2: ", workPath)
	for index, info := range fileInfos {
		fmt.Println(index, "=", info.Name())
		path := filepath.Join(newPath, info.Name())
		fmt.Println(path)
		if !info.IsDir() && info.Name() != "lua5.1.zip" {
			if file, err := os.Open(path); err == nil {
				dealFile(file)
			} else {
				fmt.Println(err)
			}

		}
	}
}

func dealFile(file *os.File) {
	bytes, _ := ioutil.ReadAll(file)
	encoder := mahonia.NewDecoder("gbk")
	fmt.Println(encoder.ConvertString(string(bytes)))
}
