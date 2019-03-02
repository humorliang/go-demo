package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

/*
Encode and Decode Strings using Base 64
编码和解码用base64
*/
func main() {
	myStr := "this is a string base64!"

	//encode
	encodedStr := base64.StdEncoding.EncodeToString([]byte(myStr))
	fmt.Println("encoded base64 string:", encodedStr)

	//decode
	deCodedStr, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decode str: %s\n", deCodedStr)
}
