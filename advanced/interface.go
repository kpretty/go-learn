package advanced

import "fmt"

// interface 是一种抽象类型，其接口值是引用类型。具体类型注重"我是谁"，接口类型注重"我能做什么"
// 接口更像时一种约束，概括了一种类型应该具备哪些方法

// Writer 接口定义
type Writer interface {
	Write([]byte) error
}

// 在 go 语言中，只要一个类型实现了和接口中定义的所有方法，且参数列表返回值一致，那就定义这个类型也是接口类型，实现了这个接口
// 实现类型面向对象中的多态

// Cat Dog 定义两个结构体
type Cat struct{}
type Dog struct{}

// Animal 定义动物的行为
// 即：只要结构体实现了 Say 和 Eat 就可以看成动物
type Animal interface {
	Say() // 叫
	Eat() // 吃
}

func (c Cat) Say() {
	fmt.Println("喵喵喵")
}

func (c Cat) Eat() {
	fmt.Println("我吃猫粮")
}

func (d Dog) Say() {
	fmt.Println("汪汪汪")
}

func (d Dog) Eat() {
	fmt.Println("我吃狗粮")
}

func Display(animal Animal) {
	animal.Say()
	animal.Eat()
}

// 值接收者与引用接收者
// 结构体实现方法时可以选择值接收者或者引用接收者
// 对于值接收者实现的接口
// 		其结构体的值类型和指针类型都可以视为接口类型，这是因为在接口调用方法时值接收者的方法会将值类型复制一份到临时变量
//		变量可以通过 & 获取其指针类型，因此变量实现了指针类型所对应的方法，也就是值类型指针类型都实现了接口
// 对于指针接收者实现的接口
// 		接口传递时只传递指针没有复制过程，因此只有接口体的指针类型可以视为接口类型，值类型不行，因为本质上其值类型没有实现接口方法

type Monkey struct{}
type Fox struct{}

type Sayer interface {
	say()
}

func (m Monkey) say() {}
func (f *Fox) say()   {}
func Say(say Sayer) {
	say.say()
}

// Cell 接口组合，类似java的interface多继承
type Cell interface {
	Sayer
	Animal
}

// 同时可以通过结构体匿名成员来组合多个结构体实现对接口的视线
// 例如 上述的 Animal 需要实现两个方法

type a1 struct{}
type a2 struct{}

func (a a1) Say() {}
func (a a2) Eat() {}

// a1 a2 都不是 Animal 类型，但可以组合两者为新的结构体
// aa 就是 Animal 类型
type aa struct {
	a1
	a2
}
