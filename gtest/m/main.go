package main

import (
	"algorithm"
	"sync"
)

func main() {
	//初始化
	fork := algorithm.NewFork()

	var wg = sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		var phor = algorithm.Philosorpher{Name: i, Fr: fork}
		wg.Add(1)
		go phor.Run()
	}
	wg.Wait()
}
