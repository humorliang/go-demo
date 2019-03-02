package car

import "fmt"

/*
工厂模式是我们最常用的实例化对象模式了，是用工厂方法代替new操作的一种模式。
工厂模式：
	简单工厂：简单工厂模式的工厂类一般是使用静态方法，通过接收的参数的不同来返回不同的对象实例。
	工厂方法：工厂方法是针对每一种产品提供一个工厂类。通过不同的工厂实例来创建不同的产品实例。
	在同一等级结构中，支持增加任意产品。
	抽象工厂：抽象工厂是应对产品族概念的。
	比如说，每个汽车公司可能要同时生产轿车，货车，客车，那么每一个工厂都要有创建轿车，货车和客车的方法。
	工厂方法需要：
		1. 工厂接口
		2. 工厂结构体
		3. 产品接口
		4. 产品结构体
	抽象工厂模式除了具有工厂方法模式的优点外，最主要的优点就是可以在类的内部对产品族进行约束。
	所谓的产品族，一般或多或少的都存在一定的关联，抽象工厂模式就可以在类内部对产品族的关联关系进行定义和描述，
	而不必专门引入一个新的类来进行管理。
*/

//工厂接口
type CarFactory interface {
	//生产小车方法
	CreateMiniCar() CarInterface
	//生产大车方法
	CreateBigCar() CarInterface
}

//车辆接口
type CarInterface interface {
	//车辆信息
	Info() string
}

//定义大车类型
type BigCar struct {
	brand string
}

//定义小车类型
type MiniCar struct {
	brand string
}

//定义上海工厂
type ShFactory struct {
}

//定义深圳工厂
type SzFactory struct {
}

//实现车辆接口
func (bc *BigCar) Info() string {
	info := fmt.Sprintf("这是%s大车的信息\n", bc.brand)
	return info

}
func (mc *MiniCar) Info() string {
	info := fmt.Sprintf("这是%s小车的信息\n", mc.brand)
	return info
}

//实现工厂接口
func (sh *ShFactory) CreateMiniCar() CarInterface {
	//这里必须要初始化一个类型指针 才能给接口
	minCar := new(MiniCar) //等同于 &MiniCar{}进行初始化  而 MiniCar{}只表示一个类型
	minCar.brand = "上海牌"
	return minCar
}
func (sh *ShFactory) CreateBigCar() CarInterface {
	bigCar := new(BigCar)
	bigCar.brand = "上海牌"
	return bigCar
}
func (sz *SzFactory) CreateMiniCar() CarInterface {
	minCar := new(MiniCar)
	minCar.brand = "深圳牌"
	return minCar
}
func (sz *SzFactory) CreateBigCar() CarInterface {
	bigCar := new(BigCar)
	bigCar.brand = "深圳牌"
	return bigCar
}
