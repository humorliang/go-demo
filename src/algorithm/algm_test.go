package algorithm

/*

要开始一个单元测试，需要准备一个 go 源码文件，在命名文件时需要让文件必须以_test结尾。

单元测试源码文件可以由多个测试用例组成，每个测试用例函数需要以Test为前缀，例如：
func TestXXX( t *testing.T )

测试用例文件不会参与正常源码编译，不会被包含到可执行文件中。
测试用例文件使用 go test 指令来执行，没有也不需要 main() 作为函数入口。所有在以_test结尾的源码内以Test开头的函数会自动被执行。
测试用例可以不传入 *testing.T 参数。
*/
import "testing"

func TestSelectionSort(t *testing.T) {
	s := []int{2, 10, 12, 3, 8, 6, 1, 4, 20, 19}
	t.Log("sort before:", s)
	SelectionSort(s)
	t.Log("sort after:", s)
}

func TestInsertionSort(t *testing.T) {
	s := []int{2, 10, 12, 3, 8, 6, 1, 4, 20, 19}
	t.Log("sort before:", s)
	InsertionSort(s)
	t.Log("sort after:", s)
}

func TestBinaryFind(t *testing.T) {
	s := []int{1, 3, 5, 7, 8, 10, 18, 19, 20, 25}
	BinaryFind(&s, 0, len(s)-1, 3)
}

func TestBubbleAscSort(t *testing.T) {
	s := []int{2, 10, 12, 3, 8, 6, 1, 4, 20, 19}
	t.Log("sort before:", s)
	BubbleAscSort(s)
	t.Log("sort after:", s)
}

func TestBubbleDescSort(t *testing.T) {
	s := []int{2, 10, 12, 3, 8, 6, 1, 4, 20, 19}
	t.Log("sort before:", s)
	BubbleDescSort(s)
	t.Log("sort after:", s)
}

func TestQuickSort(t *testing.T) {
	s := []int{2, 10, 12, 3, 8, 6, 1, 4}
	//s := []int{1, 2, 3, 4, 5, 6}
	t.Log("sort before:", s)
	QuickSort(s)
	t.Log("sort after:", s)
}

/*
基准测试——获得代码内存占用和运行效率的性能数据
基准测试可以测试一段程序的运行性能及耗费 CPU 的程度。
Go 语言中提供了基准测试框架，使用方法类似于单元测试，使用者无须准备高精度的计时器和各种分析工具，基准测试本身即可以打印出非常标准的测试报告。
$ go test -v -bench=. benchmark_test.go
2) 基准测试原理
基准测试框架对一个测试用例的默认测试时间是 1 秒。开始测试时，当以 Benchmark 开头的基准测试用例函数返回时还不到 1 秒，那么 testing.B 中的 N 值将按 1、2、5、10、20、50……递增，同时以递增后的值重新调用基准测试用例函数。
3) 自定义测试时间
通过-benchtime参数可以自定义测试时间，例如：
$ go test -v -bench=. -benchtime=5s benchmark_test.go
4) 测试内存
基准测试可以对一段代码可能存在的内存分配进行统计，下面是一段使用字符串格式化的函数，内部会进行一些分配操作。
func Benchmark_Alloc(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("%d", i)
    }
}
在命令行中添加-benchmem参数以显示内存分配情况，参见下面的指令：
$ go test -v -bench=Alloc -benchmem benchmark_test.go
*/
func BenchmarkInsertionSort(b *testing.B) {
	s := []int{2, 10, 12, 3, 8, 6, 1, 4, 20, 19}
	for i := 0; i < b.N; i++ {
		InsertionSort(s)
	}
}

func BenchmarkBubbleAscSort(b *testing.B) {
	s := []int{2, 10, 12, 3, 8, 6, 1, 4, 20, 19}
	for i := 0; i < b.N; i++ {
		BubbleAscSort(s)
	}
}
func BenchmarkSelectionSort(b *testing.B) {
	s := []int{2, 10, 12, 3, 8, 6, 1, 4, 20, 19}
	for i := 0; i < b.N; i++ {
		SelectionSort(s)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	//s := []int{2, 10, 12, 3, 8, 6, 1, 4, 20, 19}
	s := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < b.N; i++ {
		QuickSort(s)
	}
}
