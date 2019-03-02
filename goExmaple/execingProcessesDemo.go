package main

import (
	"os/exec"
	"os"
	"syscall"
	"fmt"
)

/*
用其他的（也许是非 Go 程序）来完全替代当前的 Go 进程。这时候，我们可以使用经典的 exec方法的 Go 实现。
*/
func main() {
	//获取文件可执行的绝对路径
	binary, lookErr := exec.LookPath("ls")
	fmt.Println("可执行路径：", binary)
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()
	//Exec 需要的参数是切片的形式的
	//Exec 同样需要使用环境变量。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
