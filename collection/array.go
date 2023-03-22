package collection

import (
	"fmt"
	"math"
)

// 数组，保存相同类型数据且长度固定，包含类型和长度
func arraySwap() {
	var a1 [2]int
	var a2 [3]int
	var a3 [3]int
	fmt.Printf("a1=%v\n", a1)
	fmt.Printf("a2=%v\n", a2)
	// 类型不同无法赋值
	// a1 = a2
	a3 = a2
	fmt.Printf("a2=%v\n", a3)
}

// 数组定义
func defineArray() {
	// 声明定长数组
	var a1 [10]int // 自动初始化
	fmt.Printf("a1=%v type %T\n", a1, a1)
	// 声明定长数组并初始化
	a2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a1=%v type %T\n", a2, a2)
	// 如果自己初始化，可以省略长度让编译器自行推断
	a3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("a1=%v type %T\n", a3, a3)
	// 可以指定某个索引进行初始化
	a4 := [...]int{1: 1, 2: 2, 10: 10} // 给索引位置为1,2,10位置进行初始化
	fmt.Printf("a1=%v type %T\n", a4, a4)
}

// 数组是值传递
func arrayIsValuePass(src [5]int) {
	src[3] = math.MaxInt
}

// 切片是引用传递
func sliceIsValuePass(src []int) {
	src[3] = math.MaxInt
}

// 两数之和 有序数组版
func twoNumSum(array []int, target int) [][]int {
	var result [][]int
	// 定义两个指针指向 array 头尾
	start, end := 0, len(array)-1
	for start < end {
		if array[start]+array[end] == target {
			result = append(result, []int{start, end})
			start++
		} else if array[start]+array[end] > target {
			end--
		} else {
			start++
		}
	}
	return result
}
