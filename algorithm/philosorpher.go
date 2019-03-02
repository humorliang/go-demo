package algorithm

import (
	"fmt"
	"time"
	"sync"
)

/*


有五个哲学家，他们的生活方式是交替地进行思考和进餐。他们共用一张圆桌，分别坐在五张椅子上。
在圆桌上有五个碗和五支筷子，平时一个哲学家进行思考，饥饿时便试图取用其左、右最靠近他的筷子，只有在他拿到两支筷子时才能进餐。
进餐完毕，放下筷子又继续思考。

哲学家进餐问题的改进解法
方法一：至多只允许四位哲学家同时去拿左筷子，最终能保证至少有一位哲学家能进餐，并在用完后释放两只筷子供他人使用。
方法二：仅当哲学家的左右手筷子都拿起时才允许进餐。
方法三：规定奇数号哲学家先拿左筷子再拿右筷子，而偶数号哲学家相反。


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
