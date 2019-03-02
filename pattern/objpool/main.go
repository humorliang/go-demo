package main

import (
	"sync"
	"pattern/objpool/pool"
	"fmt"
)

func main() {
	p := pool.New(3)
	wait := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		index := i
		wait.Add(1)
		go func(pool pool.Pool, ind int) {
			select {
			case i := <-*p:
				i.Do(ind)
				*p <- i
			default:
				fmt.Println("等待对象池对象")
			}
			wait.Done()
		}(*p, index)
	}
	wait.Wait()
}
