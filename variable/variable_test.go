package variable

import (
	"fmt"
	"testing"
)

// 测试变量默认值
func TestVariableDefault(t *testing.T) {
	fmt.Println("name = ", name)
	fmt.Println("age = ", age)
	fmt.Println("isOk = ", isOk)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
}

// 测试类型推导
func TestTypeInference(t *testing.T) {
	fmt.Println("_name = ", _name)
	fmt.Println("_age = ", _age)
	fmt.Println("_a = ", _a)
	fmt.Println("_b = ", _b)
}

// 测试短变量声明
func TestShortVariableDeclaration(t *testing.T) {
	shortVariableDeclaration()
}

// 测试匿名变量
func TestFoo(t *testing.T) {
	id, _ := foo()
	_, name := foo()
	fmt.Printf("id = %d, name = %s\n", id, name)
}

// 测试变量交换
func TestSwapVar(t *testing.T) {
	swapVar()
}
