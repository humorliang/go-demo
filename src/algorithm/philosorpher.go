package algorithm

import (
	"fmt"
	"time"
	"sync"
)

/*
问题描述：一圆桌前坐着5位哲学家，两个人中间有一只筷子，桌子中央有面条。
哲学家思考问题，当饿了的时候拿起左右两只筷子吃饭，必须拿到两只筷子才能吃饭。
上述问题会产生死锁的情况，当5个哲学家都拿起自己右手边的筷子，准备拿左手边的筷子时产生死锁现象。

2、每个哲学家必须确定自己左右手的筷子都可用的时候，才能同时拿起两只筷子进餐，吃完之后同时放下两只筷子。
*/
//定义哲学家
type Philosorpher struct {
	Name int
	Fr   Fork
}

//Fork
type Fork struct {
	mu   sync.Mutex   //锁
	chop [5]chan bool //筷子
}

//吃
func (this Philosorpher) Eating() {
	fmt.Println("哲学家：", this.Name, "正在吃...")
	time.Sleep(1 * time.Second)
}

//想
func (this Philosorpher) Thinking() {
	fmt.Println("哲学家：", this.Name, "正在想...")
	time.Sleep(1 * time.Second)
}

//哲学家干的事情
func (this Philosorpher) Run() {
	for true {
		this.Thinking()
		this.Fr.TakeFork(this)
		this.Eating()
		this.Fr.PutFork(this)
	}
}

//初始化筷子
func NewFork() Fork {
	var ch01 = make(chan bool, 1)
	var ch02 = make(chan bool, 1)
	var ch03 = make(chan bool, 1)
	var ch04 = make(chan bool, 1)
	var ch05 = make(chan bool, 1)
	ch01 <- true
	ch02 <- true
	ch03 <- true
	ch04 <- true
	ch05 <- true

	return Fork{
		mu:   sync.Mutex{},
		chop: [5]chan bool{ch01, ch02, ch03, ch04, ch05},
	}
}

/*
人和筷子的位置
		   0    p0     1
	p4   		            p1

	4						2

	p3			3			p2

*/

//拿筷子
func (this Fork) TakeFork(p Philosorpher) {
	this.mu.Lock()
	defer this.mu.Unlock()
	//获取哲学家
	pId := p.Name
	//当哲学家拿到一根筷子没有另外一个筷子的时候进行通道阻塞，
	select {
	//当满足一根筷子的时候进行另一个筷子判断
	case <-this.chop[pId]:
		//当这根筷子满足时吃饭，不满足进行阻塞
		select {
		case <-this.chop[(pId+1)%5]:
		}
	case <-this.chop[(pId+1)%5]:
		select {
		case <-this.chop[pId]:
		}
	}
	fmt.Println("拿筷子")
}

//放筷子
func (this Fork) PutFork(p Philosorpher) {
	pId := p.Name
	//向通道里面进行写  代表放筷子
	this.chop[pId] <- true
	this.chop[(pId+1)%5] <- true
	fmt.Println("放筷子")
}
