package main

import (
	"net/url"
	"fmt"
	"runtime"
	"os/user"
	"log"
	"os"
)

/*
How to check if a string is a URL
检查字符串是否是链接

Detect if Code is Running On Windows (at Runtime)
判断系统的环境

Get the Current Username, Name and Home Dir (Cross Platform)
跨平台的获取当前的用户信息

*/
func main() {
	//链接判断
	str1 := "http://www.baidu.com"
	fmt.Println("is url:", isVaildUrl(str1)) //true
	str2 := "www.baidu.com"
	fmt.Println("is url:", isVaildUrl(str2)) //false

	//操作系统
	fmt.Println("系统环境：", runtime.GOOS)

	//系统用户信息
	userInfo, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name:", userInfo.Name, "userId:", userInfo.Uid)
	fmt.Println("Home Dir:", userInfo.HomeDir, "userName:", userInfo.Username)
	fmt.Println("real name:",os.Getenv("SUDO_USER"))
	/*
		sudo go run main.go

		name: System Administrator
		userId: 0
		Home Dir: /var/root
		userName: root
		real name: user001
	*/


}

//判断字符串是否是链接
func isVaildUrl(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}
	return true
}
