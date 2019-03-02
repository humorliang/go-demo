package main

import "fmt"

/*
接口 是方法特征的命名集合。
interface是一组method签名的组合，我们通过interface来定义对象的一组行为
*/
type employee struct {
	name   string
	salary string
}
type student struct {
	name       string
	speciality string
}

// 定义接口
type person interface {
	sayHi() float64
	sayGo() float64
}

func (e employee) sayHi() float64 {
	fmt.Println("employee say hi")
	return 0
}
func (e employee) sayGo() float64 {
	fmt.Println("employee say go")
	return 0
}
func (s student) sayHi() float64 {
	fmt.Println("stu say hi")
	return 0
}
func (s student) sayGo() float64 {
	fmt.Println("stu say Go")
	return 0
}
func say(p person) {
	fmt.Println(p.sayGo())
	fmt.Println(p.sayHi())
}
func main() {
	e := employee{"a", "100"}
	s := student{"b", "history"}
	
	say(e)
	say(s)
}
