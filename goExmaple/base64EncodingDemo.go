package main

import (
	"encoding/base64"
	"fmt"
)

/*
base64编码：
Go 同时支持标准的和 URL 兼容的 base64 格式。编码需要使用 []byte 类型的参数，所以要将字符串转成此类型。
*/
func main() {
	data := "ab123dsd123@#'#!!?()'+=~eqw"

	//标准base64编码
	sEc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("标准base64: ",string(sEc))
	//标准base64解码
	sDec, _ := base64.StdEncoding.DecodeString(sEc)
	fmt.Println(string(sDec))

	//url兼容base64编码
	uEc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("url兼容base64: ",uEc)
	//url兼容base64解码
	uDec, _ := base64.URLEncoding.DecodeString(uEc)
	fmt.Println(string(uDec))
}
