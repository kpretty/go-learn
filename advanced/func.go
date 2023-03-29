package advanced

import (
	"fmt"
	"unsafe"
)

// Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”。
// fun 函数名([参数列表]) [返回值] {
//		函数体
// }

func sum(x int, y int) int {
	return x + y
}

// 可变参数
func rangeAdd(args ...int) int {
	// 可变参数的本质是切片
	fmt.Printf("%T\n", args)
	var result int
	for _, arg := range args {
		result += arg
	}
	return result
}

// 多返回值
func calc1(x, y int) (int, int) {
	return x + y, x - y
}

// 多返回值给返回值命名
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

// 函数也是一种类型
func funcType() {
	fmt.Printf("%T\n", calc1)
	fmt.Printf("%d\n", unsafe.Sizeof(calc1))
}

// 函数是值类型，当函数作为函数参数传递时，会复制一份
// 定义一个数据类型 operation 是一个函数类型，满足一进一出
type operation func(int) int

// 通过函数修改函数逻辑
func modify1(op operation) {
	// op 是个函数指针，通过 * 获取到实际函数，即解引用
	op = func(x int) int {
		return x * x
	}
}

// 通过函数指针修改函数逻辑
func modify2(op *operation) {
	// op 是个函数指针，通过 * 获取到实际函数，即解引用
	*op = func(x int) int {
		return x * x * x
	}
}

// 将函数通过指针的形式传递，虽然可以减少函数本身的复制
// 但是这种操作是危险的，可能会修改函数内部逻辑产生未知bug
func funcPointer() {
	// 声明并初始化一个 operation 变量
	var op operation = func(i int) int {
		return i + 1
	}
	// 调用 op
	fmt.Println(op(2)) // 3
	// 将 op 分别作为参数传递
	modify1(op)
	// 调用 op
	fmt.Println(op(2)) // 3
	modify2(&op)
	// 调用 op
	fmt.Println(op(2)) // 8
}

// 匿名函数
// 如果函数只需要使用一次，那么就不需要为其单独定义，减少命名空间的使用
func customDisplay(value interface{}, consumer func(interface{})) {
	consumer(value)
}

// 延迟函数 defer
// defer 修饰的函数会在其归属函数即将返回(结束)时被调用
// defer 会维护一个栈，多个 defer 修饰的函数遵循先入后出的逆序执行
// defer 通常用于资源的释放，在获取资源后紧接着使用 defer 修饰资源释放方法
// 即使后续程序报错，defer 也会在返回报错堆栈信息前执行，有其他语言 finally 作用
func deferDemo(flag bool) {
	defer func() {
		fmt.Println("函数一")
	}()

	defer func() {
		fmt.Println("函数二")
	}()

	if flag {
		panic("报错了")
	}

	defer func() {
		fmt.Println("函数三")
	}()
}

// 测试题
func practice() {
	calc := func(index string, a, b int) int {
		ret := a + b
		fmt.Println(index, a, b, ret)
		return ret
	}
	x := 1
	y := 2
	// calc("AA", 1, calc("A", 1, 2))
	defer calc("AA", x, calc("A", x, y))
	x = 10
	// calc("BB", 10, calc("BB", 10, 2))
	defer calc("BB", x, calc("B", x, y))
	y = 20
	// defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	// 因此被 defer 修饰的函数其参数如果有要计算的，先计算再注册到 defer
	// calc("A", 1, 2) 		-> 打印 A 1 2 3 返回3
	// calc("B", 10, 2)		-> 打印 B 10 2 12 返回12
	// calc("BB", 10, 12)	-> 打印 BB 10 12 22 返回22
	// calc("AA", 1, 3)		-> 打印 AA 1 3 4 返回4
}
