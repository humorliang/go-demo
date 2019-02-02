package main

import (
	"os"
	"log"
	"encoding/csv"
)

/*
Write data to a CSV file
书写数据到一个csv文件
*/
func main() {
	var data = [][]string{{"line1", "hello world one"}, {"line2", "hello world two"}}

	//创建csv文件对象
	file, err := os.Create("res.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//写入文件
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//文件内容写入csv
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatal(err)
		}
	}

}
