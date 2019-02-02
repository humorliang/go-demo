package main

import "pattern/strategy/strate"

func main() {
	//定义事件
	aTob := &strate.AToB{ABDistance: 500}

	//使用策略
	//使用自行车策略
	aTob.Strategy = &strate.BikeStrategy{Speed: 20}
	aTob.Do()
	//使用汽车策略
	aTob.Strategy = &strate.BusStrategy{Speed: 100}
	aTob.Do()

}
