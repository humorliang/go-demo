package obser

import "fmt"

/*
 	观察者模式（Observer Pattern）
	定义对象间的一种一对多的依赖关系,
 	当一个对象的状态发生改变时, 所有依赖于它的对象都得到通知并被自动更新。
	//类似发布订阅
*/

//定义观察者接口
type observer interface {
	Notify(interface{})
}

//发布者
type Subject struct {
	//观察者集合
	observers []observer
	//状态信息
	status string
}

//设置状态
func (s *Subject) State(state string) {
	//变换状态  发布通知
	s.status = state
	s.NotifiyAllObserver()
}

//通知所有的观察者
func (s *Subject) NotifiyAllObserver() {
	for _, ob := range s.observers {
		//通知观察者 s发布
		ob.Notify(s)
	}
}

//添加观察者
func (s *Subject) Attach(obs ...observer) {
	s.observers = append(s.observers, obs...)
}

//定义观察者A
type ObserverObj struct {
	Id string
}

//接收通知
func (oa *ObserverObj) Notify(sub interface{}) {
	//使用断言确定发布者
	fmt.Println(oa.Id, " recevie ", sub.(*Subject).status)
}
