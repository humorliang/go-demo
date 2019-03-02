package main

import (
	"fmt"
	"time"
	"log"
)

/*
Parsing Dates from a String and Formatting
解析时间字符串格式
*/
func main() {
	//时间字符串
	dateStr:="2018-10-10 12:04"
	fmt.Println("my date string:\t",dateStr)

	//时间解析
	// 2006-01-02T15:04:05+07:00
	date,err:=time.Parse("2006-01-02 15:04",dateStr)
	if err!=nil {
		log.Fatal(err)
	}
	// Another Example:
	// myDateString := "01-01-2018"
	// myDate, err := time.Parse("02-01-2006", myDateString)

	fmt.Println("date reformatted:\t",date.Format(time.RFC3339))
	//	这里必须要用这个字符串格式：2006-01-02T15:04:05+07:00
	// 	时间类型的string底层是：
	// 	s := t.Format("2006-01-02 15:04:05.999999999 -0700 MST")
	// 	所以必须要用这个字符串
	fmt.Println("date reformatted:\t",date.Format("2006-01-02 15:04:05"))
}
