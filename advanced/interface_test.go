package advanced

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDisplay(t *testing.T) {
	cat := Cat{}
	dog := Dog{}
	Display(cat)
	Display(dog)
}

func TestSay(t *testing.T) {
	// 分别获取接口体的值类型和指针类型
	monkey1 := Monkey{}
	monkey2 := &monkey1
	fox1 := Fox{}
	fox2 := &fox1
	Say(monkey1)
	Say(monkey2)
	// 下面这个编译报错
	// Say(fox1)
	Say(fox2)
}

func TestInterface(t *testing.T) {
	// 使用反射探寻接口本质
	var say Sayer
	fmt.Println("say type ", reflect.TypeOf(say))
	fmt.Println("say value ", reflect.ValueOf(say))
	say = Monkey{}
	fmt.Println("say type ", reflect.TypeOf(say))
	fmt.Println("say value ", reflect.ValueOf(say))
	say = &Monkey{}
	fmt.Println("say type ", reflect.TypeOf(say))
	fmt.Println("say value ", reflect.ValueOf(say))
}

func TestAA(t *testing.T) {
	a := aa{}
	// 如何在编译期间校验结构体是否实现了某个接口
	// var _ Sayer = aa{} // 报错
	var _ Animal = aa{} // 不报错
	Display(a)
}
