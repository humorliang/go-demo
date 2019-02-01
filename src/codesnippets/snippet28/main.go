package main

import (
	"strings"
	"fmt"
)

/*
Search and Replace in a String
搜索和取代一个字符串
*/

func main() {
	//Example 1:
	text := "this is text"
	//字符串替换，父字符串 旧子串 新子串 替换次数 -1是任何次数
	text = strings.Replace(text, "t", "T", 1)
	fmt.Println(text)

	//Example 2:
	text = "this is go go language "
	//字符串替换，父字符串 旧子串 新子串 替换次数 -1是任何次数
	text = strings.Replace(text, "go", "", 1)
	fmt.Println(text)

	//Example 3:
	text = "this is a 'apple' and banana "
	//字符串替换，父字符串 旧子串 新子串 替换次数 -1是任何次数
	text = strings.Replace(text, "'", "\\'", -1)
	fmt.Println(text)
}
