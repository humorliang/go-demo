package main

import (
	"os"
	"fmt"
	"path/filepath"
)

/*
Cross Platform File Paths
创建一个跨平台的路径
*/

func main() {
	//linux:dir/test  windows:dir\test

	path := "dir" + string(os.PathSeparator) + "test"
	fmt.Println("path1:", path)

	path = filepath.FromSlash("dir/test")
	fmt.Println("path2:",path)
}
