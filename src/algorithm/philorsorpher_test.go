package algorithm

import (
	"testing"
	"sync"
)

func TestPhilosopher(t *testing.T) {
	//初始化
	fork := NewFork()

	var wg = sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		var phor = Philosorpher{Name: i, Fr: fork}
		wg.Add(1)
		go phor.Run()
	}
	wg.Wait()

}
