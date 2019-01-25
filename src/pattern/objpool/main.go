package main

import (
	"pattern/objpool/pool"
)

func main() {
	p := pool.New(10)
	select {
	case obj := <-p:

	}
}
