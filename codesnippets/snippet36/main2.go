package main

import (
	"archive/zip"
	"path/filepath"
	"strings"
	"os"
	"fmt"
	"io"
	"sync"
	"time"
	"log"
)

/*
解压方法二：并发
在没有使用并并发的情况下大概耗时：use time :1.86925458s
使用并发的情况下大概耗时：: 232.640019ms
*/

//定义压缩文件成员信息
type fileDesp struct {
	f     *zip.File
	fInfo io.ReadCloser
	fPath string
}

func main() {
	start := time.Now()
	files, err := unZip2("test.zip", "test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unziped done:\n " + strings.Join(files, "\n"))
	fmt.Printf("use time :%s\n", time.Since(start))
}

func unZip2(src string, dest string) ([]string, error) {
	var filenames []string
	//读取压缩文件
	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	//文件信息通道
	fPathChannel := make(chan fileDesp, len(r.File))
	//文件具体信息放入通道
	for _, f := range r.File {
		//返回文件内容信息
		fInfo, err := f.Open()
		if err != nil {
			return filenames, err
		}
		//文件具体信息
		fpath := filepath.Join(dest, f.Name)
		//压缩成员信息
		fd := fileDesp{f: f, fInfo: fInfo, fPath: fpath}
		//传入通道
		fPathChannel <- fd
		filenames = append(filenames, fpath)
		//路径类型判断
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal is path")
		}
	}
	//多go解压文件
	var wg sync.WaitGroup
	for i := 1; i <= len(r.File); i++ {
		wg.Add(1)
		go func() {

			//对通道的值进行读取
			select {
			case fd := <-fPathChannel:
				if err := createFilePath(fd); err != nil {
					break
				}
			default:
			}
		}()
		wg.Done()
	}
	wg.Wait()
	return filenames, nil
}

//具体内容的写入
func createFilePath(fd fileDesp) error {
	if fd.f.FileInfo().IsDir() {
		//创建文件夹
		os.MkdirAll(fd.fPath, os.ModePerm)
	} else {
		//创建文件
		//创建文件的文件夹
		if err := os.MkdirAll(filepath.Dir(fd.fPath), os.ModePerm); err != nil {
			return err
		}
		//创建文件
		outFile, err := os.OpenFile(fd.fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}

		//数据拷贝
		_, err = io.Copy(outFile, fd.fInfo)

		//关闭文件
		outFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
