package main

import (
	"net"
	"github.com/smallnest/rpcx/log"
	"io"
	"os"
	"fmt"
)

func mustCopy(dst io.Writer, src io.Reader) error {
	wNum, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("writer byte num:", wNum)
	return nil
}

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	err = mustCopy(os.Stdout, con)
	if err != nil {
		log.Fatal(err)
	}
	//buf := bytes.Buffer{}
	//req.Read(buf.Bytes())
	//fmt.Println(buf.String())
}
