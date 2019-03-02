package main

import (
	"os"
	"log"
	"encoding/csv"
	"fmt"
)

/*
Read a CSV File into a Struct
读取一个csv文件绑定到结构体
*/
type CsvLine struct {
	Column1 string
	Column2 string
}

func main() {
	//打开文件
	filePath := "test.csv"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	//读取所有行
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	//遍历行数据进行绑定
	for _, line := range lines {
		data := CsvLine{
			Column1: line[0],
			Column2: line[1],
		}
		fmt.Println(data.Column1 + " " + data.Column2)
	}

}
