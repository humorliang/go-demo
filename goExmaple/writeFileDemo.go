package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)

/*
写入文件
*/
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//io模块
	d1 := []byte("hello golang\n")
	err := ioutil.WriteFile("./test02.txt", d1, 0644)
	checkError(err)

	//os模块
	f, err := os.Create("./test03.txt")
	checkError(err)

	defer f.Close()
	d2 := []byte{115, 162, 182, 126}
	n2, err := f.Write(d2)
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	//将缓冲区信息写入磁盘
	f.Sync()
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("this buffer string\n")
	checkError(err)
	fmt.Printf("wrote %d byte\n", n4)

	//确保所有的缓存都写入底层
	w.Flush()

}
