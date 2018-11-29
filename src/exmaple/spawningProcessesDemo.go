package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
)

/*
生成进程：
exec.Command 函数帮助我们创建一个表示这个外部进程的对象
.Output 是另一个帮助我们处理运行一个命令的常见情况的函数，
它等待命令运行完成，并收集命令的输出。如果没有出错，dateOut 将获取到日期信息的字节。
*/
func main() {
	//外部进程对象
	datecommend := exec.Command("date")
	//
	dateout, err := datecommend.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(">date")
	fmt.Println(string(dateout))

	//外部进程stdin输入数据，从stdout搜集结果
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println(grepBytes)

	//外部bash调用 使用 bash -c
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(lsOut)
}
