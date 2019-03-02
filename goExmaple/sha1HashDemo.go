package main

import (
	"crypto/sha1"
	"fmt"
)

/*
SHA1 散列经常用生成二进制文件或者文本块的短标识
*/
func main() {
	s := "sha1 is string"
	//产生一个散列值的方式
	h := sha1.New()
	//字符散列  如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(s))
	//Sum 的参数可以用来将现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//bs := h.Sum([]byte("qqw"))
	fmt.Println(s)
	fmt.Printf("%x\n", bs)//6330bfa2e686af0bbf282e0f1f0778b4562f2067

}
