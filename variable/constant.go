package variable

// 常量在定义时必须被赋值，且程序运行期间不允许修改
const pi = 3.1415926

// 和变量一样，可以批量声明
const (
	e   = 2.71828
	c10 = 0.1234567
)

// iota 是go语言的常量计数器，只能在常量的表达式中使用在const关键字出现时将被重置为0。
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
const (
	n1 = iota
	n2
	n3 = 100
	n4 = iota
)

// 出了作用于 iota 就会被重置
const n5 = iota

// 定义字节单位
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)
