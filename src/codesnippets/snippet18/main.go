package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

/*
Generating a SHA256 HMAC Hash
生成加密字符串
HMAC是密钥相关的哈希运算消息认证码，HMAC运算利用哈希算法，以一个密钥和一个消息为输入，生成一个消息摘要作为输出
*/

func main() {
	//秘钥
	secret := "secret"
	//数据
	data := "this is data"

	fmt.Println("my secret:", secret, " my data:", data)
	//生成加密对象  算法
	h := hmac.New(sha256.New, []byte(secret))

	//数据写入对象
	h.Write([]byte(data))

	//获取结果返回一个十六进制的字符串
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println("result hash:", sha)
}
