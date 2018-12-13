package main

import (
	"fmt"
	"errors"
)

/*
按照惯例，错误通常是最后一个返回值并且是 error 类型，一个内建的接口。
errors.New 构造一个使用给定的错误信息的基本error 值。
返回错误值为 nil 代表没有错误。
通过实现 Error 方法来自定义 error 类型是可以的。这里使用自定义错误类型来表示上面的参数错误。
*/

//自定义一个error类型
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	//定义错误类型的Error()方法
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

//errors.New()返回错误信息
func fun1(arg int) (int, error) {
	if arg == 20 {
		return -1, errors.New("can't work with 20")
	}
	return arg, nil
}

//
func fun2(arg int) (int, error) {
	if arg == 20 {
		return -1, &argError{arg, "it can't work"}
	}
	return arg, nil
}

func main() {
	for _, i := range []int{2, 20} {
		//判断是否错误信息
		if r, e := fun1(i); e != nil {
			fmt.Println("fun1 is faild", e)
		} else {
			fmt.Println("fun1 is worked", r)
		}

		if r, e := fun2(i); e != nil {
			fmt.Println("fun2 is faild", e)
		} else {
			fmt.Println("fun2 is worked", r)
		}
	}

	_, e := fun2(20)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
