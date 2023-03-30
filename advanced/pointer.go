package advanced

import (
	"fmt"
)

// 指针是go的一种数据类型，使用 8 位存储内存的地址数据
// go 语言的指针不能进行偏移和运算

// 指针地址和指针类型
// go 语言的值类型都有对应的指针类型，即 *T，例如：T 是 int 类型，对应的指针类型就是 *int
var a int
var ptr = &a

func pointerNature() {
	fmt.Printf("%v, %T and ptr of ptr %p\n", ptr, ptr, &ptr)
}

// 指针取值，通过对指针类型使用 * 操作，可以获取指针类型指向内存区域的值
// interface{} 在 go 中可以表示任意类型
func getPointerValue(ptr interface{}) interface{} {
	// 判断是否是指针类型
	switch ptr.(type) {
	case *int:
		return *(ptr.(*int))
	default:
		return nil
	}
}

// new 和 make
// new 		为值类型分配内存初始化为类型零值，返回指向类型的指针
// make		为slice、map、channel分配内存初始化，返回引用类型本身
// 为什么要有 new 和 make
// 对于直接声明指针类型或引用类型时默认是 nil(指针类型时值类型，函数也是值类型)
func initError() {
	// a 是一个指指针类型，只声明因此此时 a 是 nil
	var a *int
	*a = 100
	fmt.Println(*a)

	// 同理 b 也是 nil
	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)
}

func initSuccess() {
	// new(type) *type
	// 接收一个类型，返回这个类型的指针并将指向的内存地址存入该类型的零值
	var a = new(int)
	*a = 100
	fmt.Println(*a)

	// 同理 b 也是 nil
	var b = make(map[string]int)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}

/*
new 和 make 的区别
1. 二者都是用来做内存分配的。
2. make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3. 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/
