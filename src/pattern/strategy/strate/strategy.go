package strate

import "fmt"

/*
	策略模式（Strategy Pattern）
	定义一系列的算法,把它们一个个封装起来,
	并且使它们可相互替换。
	本模式使得算法可独立于 使用它的场景而变化。

	// 1.定义场景,和策略接口
	// 2.设定事件策略执行
	// 3.定义不同的的策略
*/

//定义策略接口
type Strategy interface {
	Do(interface{})
}

//定义场景
type AToB struct {
	//距离
	ABDistance float64
	//到达方式策略
	Strategy Strategy
}

//设定策略执行
func (ab *AToB) Do() {
	//判断当前场景策略
	if ab.Strategy != nil {
		//使用当前策略执行该场景
		ab.Strategy.Do(ab)
	}
}

//设定自行车策略
type BikeStrategy struct {
	Speed float64
}

func (bs *BikeStrategy) Do(ab interface{}) {
	//将ab接口类型 变成 AToB类型
	aTob, ok := ab.(*AToB)
	if ok && bs.Speed <= 0.00001 {
		return
	}
	fmt.Printf("方式：骑自行车 用时：%.2f h \n", aTob.ABDistance/bs.Speed)
}

//设定汽车策略
type BusStrategy struct {
	Speed float64
}

func (bs *BusStrategy) Do(ab interface{}) {
	aTob, ok := ab.(*AToB)
	if ok && bs.Speed <= 0.00001 {
		return
	}
	fmt.Printf("方式：坐汽车 用时：%.2f h \n", aTob.ABDistance/bs.Speed)
}
