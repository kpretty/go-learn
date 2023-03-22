package variable

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
整型
	int8		有符号8位整数(一个字节)
	int16		有符号16位整数(两个字节)
	int32		有符号32位整数(三个字节)
	int64		有符号64位整数(四个字节)
	int			32位系统为int32,64位系统为int64
	uint8		无符号8位整数(一个字节)
	uint16		无符号16位整数(两个字节)
	uint32		无符号32位整数(三个字节)
	uint64		无符号64位整数(四个字节)
	uint		32位系统为uint32,64位系统为uint64
整型默认值为 int
*/
var (
	i1  int8
	i2  int16
	i3  int32
	i4  int64
	i5  int
	i6  uint8
	i7  uint16
	i8  uint32
	i9  uint64
	i10 uint
	i11 = 0
)

// 数字字面量，可以直接定义二进制、八进制、十进制、十六进制数字
// 十进制
var v1 = 11

// 八进制，以 0 开头
var v2 = 011

// 十六进制，以 0x 开头
var v3 = 0x11

// 允许用 _ 分割数据
var v4 = 123_456

/*
浮点数
float32	最大值math.MaxFloat32
float64 最大值math.MaxFloat64
*/
var f1 = 3.1415926
var f2 = 3.1415926
var f3 = math.MaxFloat32
var f4 = math.MaxFloat64

/*
复数
complex64	实部和虚部为32位
complex128	实部和虚部为64为
*/
var c1 = 1 + 1i
var c2 = 2 + 2i

/*
布尔值
	1. 布尔类型变量的默认值为false
	2. 不允许将整型强制转换为布尔型
	3. 布尔型无法参与数值运算，也无法与其他类型进行转换
*/

/*
字符型，用 '' 表示，友有两个类型
	1. byte，本质上是uint8，代表一个 ASCII 码字符
	2. rune，本质上是int32，代表一个 Unicode 字符
默认 rune
*/
var char1 byte = 'a'
var char2 = '中'

/*
字符串，本质上是个字符数组（字节数组），但字符串是作为 go 的原生类型
*/
var s1 = "hello,中国"

// 验证 string 是个字节数组
func verifyStringIsByteArray() {
	s1 := "hello,中国"
	fmt.Println(len(s1)) // 12: 英文是 ASCII 占一个字节，中文是 Unicode 在 UTF8 编码中占三个字节
	for i := 0; i < len(s1); i++ {
		fmt.Printf("index [%d] value [%v]\n", i, s1[i])
	}
	// 对于纯英文的字符串可以使用 fori 遍历，但是包含中文这种操作是非法的，且通过下标获取字符对于 string 也是非法的
	fmt.Printf("非法获取第七个字符 %c\n", s1[6])
	// 如果需要按照常理一样获取字符串每个元素，需要将 string 转换为 rune 数组
	s2 := []rune(s1)
	fmt.Printf("正确获取第七个字符 %c\n", s2[6])
	// 都可以使用 range 进行遍历
	for i, v := range s1 {
		fmt.Printf("s1[%d]=%c\n", i, v)
	}
}

// 字符串是不可变数组，若需要修改字符串，则需要先转换成 []byte 或 []rune 再转换为 string
// 但无论转换成哪种类型数组都是重新分配内存并复制字节数组
func editString() {
	s1 := "你好！我是男生"
	s2 := []rune(s1)
	s2[5] = '女'
	s3 := string(s2)
	fmt.Printf("s1 value %s address %p\n", s1, &s1)
	fmt.Printf("s3 value %s address %p\n", s3, &s3)
}

// 强制类型转换，go 没有隐式转换
func typeTransform() {
	a := 1
	fmt.Println(float64(a))
	s := "1"
	i, _ := strconv.ParseInt(s, 10, 64)
	fmt.Printf("i = %d and type [%T]\n", i, i)
}

// 字符串拼接
func stringsConnection() {
	s1 := "你好！"
	s2 := "我是男生"
	s := s1 + s2
	fmt.Println(s)
	// + 拼接效率不高
	builder := strings.Builder{}
	builder.WriteString(s1)
	builder.WriteString(s2)
	fmt.Println(builder.String())
}
