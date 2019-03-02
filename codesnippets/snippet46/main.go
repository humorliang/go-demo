package main

import (
	"encoding/json"
	"log"
	"fmt"
)

/*
JSON Encode an Array of Objects
数组编码成json
JSON Decode into Objects
对json进行解码
*/

type Page struct {
	//字段要使用大写字母，否则编码成json无法识别
	Title    string `json:"title"`
	FileName string `json:"file_name"`
	Content  string `json:"content"`
}

var Pages = []Page{
	Page{
		"page1",
		"page1.txt",
		"this is page1 content",
	},
	Page{
		"page2",
		"page2.txt",
		"this is page2 content",
	},
}
var rowJson = []byte(`[{"title":"page1","file_name":"page1.txt","content":"this is page1 content"},
					{"title":"page2","file_name":"page2.txt","content":"this is page2 content"}]`)

func main() {
	//encode json
	encodeJson, err := json.Marshal(Pages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("encodeJson:%s \n", encodeJson)

	//decode json
	var pages []Page
	err = json.Unmarshal(rowJson, &pages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decode json:%s \n", pages)

}
