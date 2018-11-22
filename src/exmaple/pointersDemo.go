package main

import "fmt"

/*
传值和传指针
1. 当函数传递值时只会 copy一份数值，并不会改变数值本身
2. 当函数传递指针传递的是值得地址，当采用 *p 解引用时 再操作 会改变变量的地址  从而改变值
		&x  可以取到变量x地址
变量：
普通变量a其实是语言本身创造了，是为了更方便的表示内存。我们对a进行访问其实就是直接对内存进行访问。
至于a表示的内存的地址是多少，我们一般不用关心。
编译器会自动分配地址，也就是常说的为a分配一个地址。如果想知道a的地址也可以通过 &a 得知。
指针：
	1. 定义一个指针变量
		 var a int
		 ptr:= *int
		 ptr2:=new(int)
	2. 把变量地址赋值给指针
		ptr:=&a
	3. 访问指针变量中可用地址的值
		value=*ptr
*/

func sendVar(a int) {
	fmt.Println("数值操作")
	a = a + 1
}
func sendPointer(a *int) {

	//*a 将变量进行解引用
	fmt.Println(a)  //0xc000092000
	fmt.Println(*a) //10
	*a = *a + 1
}

func ptrOpt() {
	a := 12
	//指针定义1
	var ptr *int
	ptr2 := &a
	//指针定义2
	ptr3 := new(int)
	ptr3 = &a
	fmt.Println("ptr", ptr)   //nil
	fmt.Println("ptr2", ptr2) //0xc00006e040
	fmt.Println("ptr3", ptr3) //0xc00006e040

	//取指针地址
	fmt.Println(ptr2)
	//取指针指向的值 *指针
	fmt.Println(*ptr2) //12

}

func main() {
	i := 10
	fmt.Println("初始值i:", i)
	fmt.Println("初始值i地址:", &i)
	sendVar(i)
	fmt.Println("传值后i:", i)
	fmt.Println("传值后i地址:", &i)
	sendPointer(&i) //参数为指针 需要传递地址
	fmt.Println("传指针i:", i)
	fmt.Println("传指针i地址:", &i)
	fmt.Println("--------------")
	ptrOpt()

}
