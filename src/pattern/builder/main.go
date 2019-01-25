package main

import "pattern/builder/product"

func main() {
	//使用建造者模式
	//构建小卖铺
	dir := new(product.Director)
	//构建泡面商品建造者
	noodleBuilder := new(product.NoodlesBuilder)
	//商品放入小卖铺  noodle
	dir.SetBuilder(noodleBuilder)
	//得到商品
	noodles := dir.Generate()
	//显示商品信息
	noodles.ShowInfo()
}
