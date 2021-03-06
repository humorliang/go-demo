package main

import (
	"fmt"
	"strings"
)

/*
字符串函数：
标准库的 strings 包提供了很多有用的字符串相关的函数
*/

func main() {
	//
	var p = fmt.Println
	p("Contains: ", strings.Contains("tests", "es"))      //Contains:  true
	p("count: ", strings.Count("tests", "t"))             //count:  2
	p("hasPrefix: ", strings.HasPrefix("tests", "te"))    //hasPrefix:  true
	p("hasSuffix: ", strings.HasSuffix("tests", "st"))    //hasSuffix:  true
	p("index: ", strings.Index("tests", "e"))             //index:  1
	p("join: ", strings.Join([]string{"a", "b"}, "-"))   //join:  a-b
	p("Repeat: ", strings.Repeat("a", 5))                //Repeat:  aaaaa
	p("Replace: ", strings.Replace("foo", "o", "a", -1)) //Replace:  faa
	p("Replace: ", strings.Replace("foo", "o", "a", 1))  //Replace:  fao
	p("Split: ", strings.Split("a-b-c-s-d-d", "-"))      //Split:  [a b c s d d]
	p("ToLower: ", strings.ToLower("TEST"))              //ToLower:  tests
	p("ToUpper: ", strings.ToUpper("tests"))              //ToUpper:  TEST
	p()
	p("len: ", len("ads"))  //len:  3
	p("char: ", "hello"[1]) //char 101

}
