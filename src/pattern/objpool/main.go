package main

import (
	"pattern/objpool/pool"
	"sync"
)

func main() {
	p := pool.New(10)
	var wg sync.WaitGroup
	for{
		wg.Add(1)
		go func() {
			select {
			case <-p:

			}
		}()
	}

}
