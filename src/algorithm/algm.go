package algorithm

import "fmt"

/*
1.选择排序 时间复杂度：n的平方 空间复杂度：1

选择排序的基本思想是：如果有N个元素需要排序，那么首先从N个元素中找到最小的那个元素与第0位置上的元素交换
（说明一点，如果没有比原本在第0位置上的元素小的就不用交换了，后面的同样是），
然后再从剩下的N-1个元素中找到最小的元素与第1位置上的元素交换，
之后再从剩下的N-2个元素中找到最小的元素与第2位置上的元素交换，…直到所有元素都排序好（
也就是直到从剩下的2个元素中找到最小的元素与第N-2位置上的元素交换）。
*/
//选择排序
func SelectionSort(s []int) {
	l := len(s)     //排序元素个数  全局变量
	m := len(s) - 1 //控制数据索引
	//控制所有元素的访问
	for i := 0; i < m; i++ {
		//保存此时对比元素i位置
		k := i
		//从 i+1 查找最小的数 与i 进行位置交换
		for j := i + 1; j < l; j++ {
			if s[j] < s[k] {
				k = j
			}
		}
		// 判断此时k值是否改变  改变就交换值
		if k != i {
			s[k], s[i] = s[i], s[k]
		}
	}
}

/*
2.插入排序 时间复杂度：O(n^2)，空间复杂度：1

插入排序的思想就是：从数组的下标为0的元素出发，每次向后取一个元素，将该元素插入到前面已排好的子数组中，
以排升序为例，将所要插入的元素插在左边小于该元素和右边大于该元素之间的位置，进而形成新的子数组，
直到所有元素全插进来为止。
*/

//插入排序
func InsertionSort(s []int) {
	//go 语言中的切片是一个引用
	n := len(s) //元素长度
	//数据不够直接返回
	if n < 2 {
		return
	}
	//从第二个数开始每次取一个数
	for i := 1; i < n; i++ {
		//对元素进行插入
		//对已经排好的切片进行判断插入
		for j := i; j > 0; j-- {
			//采用升序判断  如果后面比前面的小则进行交换
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
}

/*
3. 二分查找：首先了解二分查找，首先在长度为n的表范围中查找，第一次循环在n/2中查找，第二次在n/2/2中查找，依次循环。假设在第X次找到，
那么 就是找2的X次方次，有2的X次方=n解出x为log2的n ，故时间复杂度为log2N。由于辅助空间是常数级别的所以：空间复杂度是O(1);
思路：先找到中间的下标middle = (leftIndex + RightIndex) /2 ,然后让中间的下标值和FindVal比较
a:如果arr[middle] > FindVal,那么就向LeftIndex~(middle - 1)区间找
b:如果arr[middle] < FindVal,那么就向middle + 1 ~RightIndex区间找
c:如果arr[middle] == FindVal,那么直接返回
②从①的a、b、c递归执行，知道找到位置
③如果LeftIndex > RightIndex，则表示找不到，退出
*/

//二分查找
//二分查找函数 //假设有序数组的顺序是从小到大（很关键，决定左右方向）
func BinaryFind(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	//判断索引
	if leftIndex > rightIndex {
		fmt.Println("not find number ")
		return
	}
	//先找到中间下标
	middle := (leftIndex + rightIndex) / 2

	//数值比较
	// 如果比中间值小则在左半区
	if (*arr)[middle] > findVal {
		//进行递归操作继续查找
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Println("finded num index is:", middle)
	}
}

/*
4.冒泡排序 时间复杂度：n平方，空间复杂度：1

冒泡排序顾名思义就是整个过程像气泡一样往上升，单向冒泡排序的基本思想是（假设由小到大排序）：
对于给定n个记录，从第一个记录开始依次对相邻的两个记录进行比较，当前面的记录大于后面的记录时，交换位置，
进行一轮比较和换位后，n个记录的最大记录将位于第n位，然后对前（n-1）个记录进行第二轮比较；重复该过程，
直到记录剩下一个为止。
*/

// 冒泡排序 升序
func BubbleAscSort(s []int) {
	l := len(s)
	//控制轮数
	for i := 0; i < l-1; i++ {
		//从每一轮剩下的数中进行产出最大数
		for j := i; j < l; j++ {
			//因为升序  若前面的大于后面的则进行数据交换
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

// 冒泡排序 降序
func BubbleDescSort(s []int) {
	l := len(s)
	//控制轮数
	for i := 0; i < l-1; i++ {
		//从每一轮剩下的数中进行产出最大数
		for j := i; j < l; j++ {
			//因为升序  若前面的小于后面的则进行数据交换
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

/*
5.快速排序：时间复杂度nlogN,最坏是n平方
它的核心思想就是选取一个基准元素(通常以需要排序的数组第一个数)，
然后通过一趟排序将 比 基准数 大 的放在右边，比基准数小的放在左边，接着对划分好的两个数组再进行上述的排序
原理：
递归+分治
每次都可以固定一个元素的最终的位置,
1. 把一个元素的最终位置给确定了,
2. 这样子该最终位置的元素的两边各自有一个数组,
3. 分别对两边的数组进行重复上面的操作,
4. 当只有一个元素的数组的时候,那么就算该位置结束了

数组拆分
*/
func QuickSort(data []int) {
	//判断数据的长度
	if len(data) <= 1 {
		return
	}
	//取到基准元素
	mid := data[0]
	//元素头尾下标
	head, tail := 0, len(data)-1

	//剩下元素找位置的过程
	//元素遍历 放置基准元素的位置
	for i := 1; i <= tail; {
		fmt.Println(data)
		/*
		[1 2 3 4 5 6] 原始
		[1 6 3 4 5 2] 第1次  data[i]=2  tail 5
		[1 5 3 4 6 2] 第2次  data[i]=6  tail 4
		[1 4 3 5 6 2] 第3次  data[i]=5  tail 3
		[1 3 4 5 6 2] 第4次  data[i]=4  tail 2

		[3 4 5 6 2]
		[3 2 5 6 4]
		[2 3 5 6 4]
		[2 3 6 5 4]

		[6 5 4]
		[5 6 4]

		[5 4]
		*/
		//如果元素 大则放右边
		if data[i] > mid {
			//将元素i 和最右边的元素进行位置互换
			data[i], data[tail] = data[tail], data[i]
			//最高下标减一
			tail--
		} else {
			//将头元素的位置 与元素i互换
			data[i], data[head] = data[head], data[i]
			//头加一
			head++
			//元素下标增加
			i++
		}
	}
	//将基准元素放到数据头
	data[head] = mid
	//递归进行双路排序
	//切片两份数据
	QuickSort(data[:head])
	QuickSort(data[head+1:])
}

/*


*/