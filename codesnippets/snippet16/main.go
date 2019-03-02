package main

import "fmt"

func main() {
	mystr := "这 个 this is a string !"

	// byte 等同于int8, 常用来处理ascii字符
	// rune 等同于int32, 常用来处理unicode或utf-8字符
	// 将string转化为rune
	// golang中string底层是通过byte数组实现的。
	// 中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。
	a := []rune(mystr) //rune  -> int32  处理unicode编码

	//进行切割
	subStr := a[0:5]
	fmt.Println(subStr)         //[36825 32 20010 32 116]
	fmt.Println(string(subStr)) //这 个 t
}
