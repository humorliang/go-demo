/*
建造者模式：是将一个复杂的对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

建造者模式通常包括下面几个角色：

1. builder：给出一个抽象接口，以规范产品对象的各个组成成分的建造。这个接口规定要实现复杂对象的哪些部分的创建，并不涉及具体的对象部件的创建。

2. ConcreteBuilder：实现Builder接口，针对不同的商业逻辑，具体化复杂对象的各部分的创建。 在建造过程完成后，提供产品的实例。

3. Director：调用具体建造者来创建复杂对象的各个部分，在指导者中不涉及具体产品的信息，只负责保证对象各部分完整创建或按某种顺序创建。

4. Product：要创建的复杂对象。
*/
package product

import (
	"fmt"
)

/*
GO的建造者模式设计步骤：
	1.产品-->2.针对产品创建抽象建造者-->3.具体化建造者(创建一类建造者，并实现抽象)
	-->4.抽象建造者的指挥者创建-->5.指挥者返回具体产品实例
使用建造者：
	1.创建指挥者--> 2.创建具体的建造者-->3.指挥者添加建造者-->4.获得建造者并使用
*/
//产品
type Product struct {
	id    int
	title string
	price float64
}

func (p *Product) ShowInfo() {
	fmt.Println("商品sku:", p.id)
	fmt.Println("商品标题:", p.title)
	fmt.Println("商品价格:", p.price)
}

//抽象建造者
type Builder interface {
	BuildId()          //建造ID
	BuildTitle()       //建造title
	BuildPrice()       //建造价格
	Bind() interface{} //返回建好的产品
}

//泡面建造者
type NoodlesBuilder struct {
	p *Product
}

func (builder *NoodlesBuilder) BuildId() {
	if builder.p == nil {
		builder.p = new(Product)
	}
	builder.p.id = 1
}

func (builder *NoodlesBuilder) BuildTitle() {
	if builder.p == nil {
		builder.p = new(Product)
	}
	builder.p.title = "这是泡面"
}

func (builder *NoodlesBuilder) BuildPrice() {
	if builder.p == nil {
		builder.p = new(Product)
	}
	builder.p.price = 30
}

func (builder *NoodlesBuilder) Bind() interface{} {
	if builder.p == nil {
		builder.p = new(Product)
	}
	return builder.p
}

//小卖铺(管理商品的指挥者)
type Director struct {
	builder Builder
}

//将要卖的商品放入小卖铺 (构建)
func (dir *Director) SetBuilder(builder Builder) {
	dir.builder = builder
}

//小卖铺和商品准备好了，开始卖东西
func (dir *Director) Generate() *Product {
	dir.builder.BuildPrice()
	dir.builder.BuildId()
	dir.builder.BuildTitle()
	//对空接口类型进行转化
	return dir.builder.Bind().(*Product)
}
