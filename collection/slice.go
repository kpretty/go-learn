package collection

import "fmt"

func sliceDemo1() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice1: value %v\t address %p\t size %d\t cap %d\n", slice1, slice1, len(slice1), cap(slice1))
	// 触发扩容
	slice2 := append(slice1, 6)
	// 扩容会创建一个新的
	fmt.Printf("slice1: value %v\t address %p\t size %d\t cap %d\n", slice1, slice1, len(slice1), cap(slice1))
	fmt.Printf("slice1: value %v\t address %p\t size %d\t cap %d\n", slice2, slice2, len(slice2), cap(slice2))
}

// 切片：拥有相同类型数据的可变长度序列，基于数组类型做的封装，支持自动扩容
// 切片定义
var s1 []string
var s2 []int
var s3 = []bool{true, false}

// 切片是引用类型，唯一合法的比较就是与 nil
// 切片本质是：一个指向底层数组的指针、长度 len 和 容量 cap
func sliceDemo2() {
	fmt.Println(s1 == nil) // 只声明，没有初始化
	fmt.Println(s2 == nil) // 只声明，没有初始化
	fmt.Println(s3 == nil)
	fmt.Printf("s3 pointer %p, value %v, len %d, cap %d\n", s3, s3, len(s3), cap(s3))
}

// 如何获取一个切片
func getSlice() {
	// 按照切片的声明方式获取切片
	s1 := []int{1, 2, 3, 4, 5} // [1 2 3 4 5] len 5 cap 5
	fmt.Printf("slice1: value %v\t address %p\t len %d\t cap %d\n", s1, s1, len(s1), cap(s1))
	// 从数组中获取切片
	a1 := [5]int{1, 2, 3, 4, 5}
	s2 := a1[:] // [start:end] => [start, end) | start 为 0 可以省略，end 为 len 可以省略
	fmt.Printf("slice1: value %v\t address %p\t len %d\t cap %d\n", s2, s2, len(s2), cap(s2))
	// 上述方式都不方便
	// 使用 make 创建切片
	// 对于使用 make 创建切片时，参数 make(切片类型, len, cap)
	s3 := make([]int, 5, 10)
	fmt.Printf("slice1: value %v\t address %p\t len %d\t cap %d\n", s3, s3, len(s3), cap(s3))
}

// 探寻切片本质
func sliceEssence() {
	// 切片本身是个引用类型，底层数据存储使用数组，切片的地址指向底层数组的起始位置
	s1 := make([]int, 5)
	s2 := s1[:2]
	s3 := s1
	fmt.Printf("slice1: value %v\t address %p\t len %d\t cap %d\n", s1, s1, len(s1), cap(s1))
	fmt.Printf("slice1: value %v\t address %p\t len %d\t cap %d\n", s2, s2, len(s2), cap(s2))
	// 切片的拷贝复制是 浅拷贝
	s3[0] = -1
	fmt.Printf("slice1: value %v\t address %p\t len %d\t cap %d\n", s1, s1, len(s1), cap(s1))
	fmt.Printf("slice1: value %v\t address %p\t len %d\t cap %d\n", s2, s2, len(s2), cap(s2))
	fmt.Printf("slice1: value %v\t address %p\t len %d\t cap %d\n", s3, s3, len(s3), cap(s3))
}

// 向切片添加元素
func appendSlice() {
	s1 := make([]int, 0)
	fmt.Printf("slice1: value %v\t address %p\t size %d\t cap %d\n", s1, s1, len(s1), cap(s1))
	// 为什么append有返回值
	// 因为切片存在扩容机制，调用 append 可能会返回原切片指向的底层数组
	// 若触发了扩容则指向新的底层数组，因此有返回值
	s1 = append(s1, 0)
	fmt.Printf("slice1: value %v\t address %p\t size %d\t cap %d\n", s1, s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	fmt.Printf("slice1: value %v\t address %p\t size %d\t cap %d\n", s1, s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	fmt.Printf("slice1: value %v\t address %p\t size %d\t cap %d\n", s1, s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	fmt.Printf("slice1: value %v\t address %p\t size %d\t cap %d\n", s1, s1, len(s1), cap(s1))
	// append 可以帮助我们做初始化，即可以给 append 传入一个未赋值的切片
	var s2 []int
	// append 可以一次插入多个元素
	s2 = append(s2, 0, 1, 2, 3, 4, 5)
	fmt.Printf("slice1: value %v\t address %p\t size %d\t cap %d\n", s2, s2, len(s2), cap(s2))
}

// 切片复制
func copySlice() {
	// 直接复制是浅拷贝，使用内置函数 copy 实现深拷贝
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5)
	// copy 需要切片已经被初始化，否则无法拷贝，注意与 append 区分
	copy(s2, s1)
	s2[0] = 0
	fmt.Printf("slice1: value %v\t address %p\t len %d\t cap %d\n", s1, s1, len(s1), cap(s1))
	fmt.Printf("slice1: value %v\t address %p\t len %d\t cap %d\n", s2, s2, len(s2), cap(s2))
}

// 删除切片元素
func deleteSlice() {
	// go 没有内置删除切片的方法，但是可以通过对切片进行切片来实现
	s1 := []int{1, 2, 3, 4, 5}
	// 删除索引位置为 2 的
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("slice1: value %v\t address %p\t len %d\t cap %d\n", s1, s1, len(s1), cap(s1))
}

// 实现删除切片
func myDelete(src []int, index int) []int {
	return append(src[:index], src[index+1:]...)
}
