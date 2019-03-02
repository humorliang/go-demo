package main

import (
	"encoding/json"
	"fmt"
)

/*

Go 提供内置的 JSON 编解码支持，包括内置或者自定义类型与 JSON 数据之间的转化。
*/
//编码
type response1 struct {
	page   int
	fruits []string
}

//解码
type response2 struct {
	page   int      `json:"page"`
	fruits []string `json:"fruits"`
}

func main() {
	//基本数据编码成json字符串
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("strings")
	fmt.Println(string(strB))

	slicD := []string{"apple", "banana", "pear"}
	slicB, _ := json.Marshal(slicD)
	fmt.Println(string(slicB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		page:   1,
		fruits: []string{"apple", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))


	//Go普通数据结构解码例子
	byt := []byte(`{"num":3.2,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

}
