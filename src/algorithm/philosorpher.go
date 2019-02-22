package algorithm

import (
	"fmt"
	"time"
	"sync"
)

//定义哲学家
type Philosorpher struct {
	Name int
	Fr   Fork
}

//Fork
type Fork struct {
	mu   sync.Mutex
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

func (this Philosorpher) Run() {
	for true {
		this.Thinking()
		this.Fr.TakeFork(this)
		this.Eating()
		this.Fr.PutFork(this)
	}
}
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

//拿筷子
func (this Fork) TakeFork(p Philosorpher) {
	this.mu.Lock()
	defer this.mu.Unlock()
	//获取哲学家
	pId := p.Name
	select {
	case <-this.chop[pId]:
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
	this.chop[pId] <- true
	this.chop[(pId+1)%5] <- true
	fmt.Println("放筷子")
}
