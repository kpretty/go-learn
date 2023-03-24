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
