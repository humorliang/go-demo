package main

/*
Print a Variable's Type (e.g. Int, String, Float)
获取变量的类型
*/
import (
	"fmt"
	"reflect"
)

func main() {
	mInt := 123
	fmt.Printf("mInt type:%T \n", mInt)

	mFloat := 12.3
	fmt.Printf("mFlaot type:%T \n", mFloat)

	mString := "this is"
	fmt.Printf("mString type:%T \n", mString)

	//使用反射进行类型获取
	mReInt := 22
	fmt.Println(reflect.TypeOf(mReInt).String())
}
