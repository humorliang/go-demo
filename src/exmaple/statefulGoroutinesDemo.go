package main

/*
另一个选择是使用内置的 Go协程和通道的的同步特性来达到同样的效果。
这个基于通道的方法和 Go 通过通信以及 每个 Go 协程间通过通讯来共享内存，
确保每块数据有单独的 Go 协程所有的思路是一致的。

state 将被一个单独的 Go 协程拥有。这就能够保证数据在并行读取时不会混乱。
*/

//定义两个读写类型
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	//
}
