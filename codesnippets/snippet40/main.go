package main

import (
	"regexp"
	"log"
	"fmt"
)

/*
Remove all Non-Alphanumeric Characters from a String (with help from regexp)
移除所有字符非字母和数字的所有字符
*/
func main() {
	example := "#sda20!ww'da"
	//创建一个正则表达式模板
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	//移除string
	processString := reg.ReplaceAllString(example, "")
	fmt.Printf("A string of %s become %s \n", example, processString)
}
