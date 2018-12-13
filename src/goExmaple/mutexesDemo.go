package main

import (
	"sync"
	"math/rand"
	"sync/atomic"
	"runtime"
	"time"
	"fmt"
)

/*
对于更加复杂的情况，我们可以使用一个互斥锁来在 Go 协程间安全的访问数据。
Lock() 这个 mutex 来确保对 state 的独占访问，读取选定的键的值，Unlock() 这个mutex，
*/
func main() {
	//共享数据
	var stat = make(map[int]int)
	//创建一把互斥锁指针
	var mutex = &sync.Mutex{}
	//计数器
	var ops uint64 = 0

	//创建读进程
	for r := 0; r < 50; r++ {
		go func() {
			total := 0
			for {
				//随机键
				key := rand.Intn(5)
				//上锁
				mutex.Lock()
				total += stat[key]
				//fmt.Println(total)
				//解锁
				mutex.Unlock()
				//原子计数
				atomic.AddUint64(&ops, 1)

				//避免饿死其他Go协程，手动释放调度片
				runtime.Gosched()
			}
		}()
	}

	//创建写进程
	for w := 0; w < 10; w++ {
		go func() {
			for {
				//随机数int [0,n)
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				stat[key] = value
				mutex.Unlock()
				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	//
	time.Sleep(time.Second * 2)
	//查看最终计数
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops", opsFinal)

	//最终锁
	mutex.Lock()
	fmt.Println("stat:", stat)
	mutex.Unlock()

}
