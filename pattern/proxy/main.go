package main

import (
	"pattern/proxy/proxy"
	"fmt"
)

func main() {
	//定义接口  事件
	var devi proxy.Device

	fmt.Println("-------------admin----------------")
	//接口赋值
	devi = &proxy.HardDiskProxy{OpId: "admin", Hd: &proxy.HardDisk{}}
	err := devi.Writer([]byte("this is a msg "))
	fmt.Println(err)
	data, _ := devi.Read()
	fmt.Printf("admin: %s \n", data)

	fmt.Println("-------------reader----------------")
	//接口赋值
	devi = &proxy.HardDiskProxy{OpId: "reader", Hd: &proxy.HardDisk{Storage: []byte("only read")}}
	err = devi.Writer([]byte("this is 2 msg "))
	fmt.Println(err)
	data, err = devi.Read()
	fmt.Println("reader:", string(data), err)

	fmt.Println("-------------writer----------------")
	//接口赋值
	devi = &proxy.HardDiskProxy{OpId: "writer", Hd: &proxy.HardDisk{}}
	err = devi.Writer([]byte("this is 3 msg "))
	fmt.Println(err)
	data, err = devi.Read()
	fmt.Println("writer:", string(data), err)
}
