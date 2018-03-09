package main

import (
	"flag"
	"github.com/fighterlyt/file2go/compress"
)

var (
	file           = ""
	packageName    = ""
	targetFileName = "file2go.go"
)

func init() {
	flag.StringVar(&file,"file", file, "数据文件名")
	flag.StringVar(&packageName,"package", packageName, "包名")
	flag.StringVar(&targetFileName,"target", targetFileName, "目标文件名")
}

func main() {
	flag.Parse()
	if err := compress.NewModel(file, packageName, targetFileName); err != nil {
		panic(err.Error())
	}
}
