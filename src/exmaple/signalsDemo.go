package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

/*
我们希望 Go 能智能的处理 Unix 信号。
例如，我们希望当服务器接收到一个 SIGTERM 信号时能够自动关机，或者一个命令行工具在接收到一个 SIGINT 信号时停止处理输入信息。
Go 通过向一个通道发送 os.Signal 值来进行信号通知
*/
func main() {
	//信号通道
	sigs := make(chan os.Signal, 1)
	done:=make(chan bool,1)

	//注册这个通过，接受特定的信号
	signal.Notify(sigs,syscall.SIGINT,syscall.SIGEMT)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
