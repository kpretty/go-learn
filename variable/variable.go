package variable

import "fmt"

/*
标识符与关键字
	标识符由字母、数字和下划线组成，并且只能以字母和下划线开头
	25个关键字
		break        default      func         interface    select
    	case         defer        go           map          struct
    	chan         else         goto         package      switch
    	const        fallthrough  if           range        type
    	continue     for          import       return       var
	37个保留字
		Constants:	true  false  iota  nil

        Types:    	int  int8  int16  int32  int64
                  	uint  uint8  uint16  uint32  uint64  uintptr
                  	float32  float64  complex128  complex64
                  	bool  byte  rune  string  error

    	Functions:	make  len  cap  new  append  copy  close  delete
					complex  real  imag
					panic  recover
*/
// 标准声明
var name string
var age int
var isOk bool

// 批量声明
var (
	a string
	b int
	c bool
	d float64
)

// 声明变量会自动对变量对应的内存区域进行初始化操作，每个变量会被初始化成其类型的默认值

// 类型推断，根据右值推断类型进行初始化
var _name = "张三"
var _age = 18

// 短变量声明，仅在函数内部使用
func shortVariableDeclaration() {
	n := 1
	m := "李四"
	fmt.Printf("n = %d, n type is [%T]\n", n, n)
	fmt.Printf("n = %s, n type is [%T]\n", m, m)
}

// 匿名变量，用 _ 表示，只占位，不占用命令空间，不分配内存，所以不存在重复声明
func foo() (int, string) {
	return 1, "王五"
}

// 交换变量
func swapVar() {
	a := 1
	b := 2
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)
}
