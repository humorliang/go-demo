package main

import "fmt"

/*
Passing Data between Go Routines with Channels
通过通道进行数据传输
*/

func main() {
	message := make(chan string)
	go work(message)

	//写数据
	msg := <-message
	fmt.Println(msg)

}

//往通道里面写数据 写数据的通道
func work(message chan<- string) {
	message <- "done"
}
