package main

import (
	"os/exec"
	"log"
	"fmt"
)

/*
Run System Commands & Binary Files
执行系统命令，或者二进制文件
*/
func main() {
	//使用os/exec库进行执行命令或二进制文件
	outContent, err := exec.Command("ls", "-alh").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", outContent)
	outContent2, err := exec.Command("pwd").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", outContent2)
}
