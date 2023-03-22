package basic

import "fmt"

// if else
func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

// if条件判断特殊写法
func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

//上下两种写法的区别在于 score 的作用域

// for
func forDemo1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

// 类似 while, go 没有 while
func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

// for range 可以遍历数组、切片、字符串、map和channel
// 数组、切片、字符串返回索引和值
// map 返回键和值
// channel 返回通道内值

// case 支持单值、多值、表达式
func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

// goto
// 如何优雅的退出双层 for
func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 2 {
				breakFlag = true
				break
			}
			fmt.Printf("i = %d, j = %d \n", i, j)
		}
		if breakFlag {
			break
		}
	}
}

// goto 慎用
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 2 {
				goto breakTag
			}
			fmt.Printf("i = %d, j = %d \n", i, j)
		}
	}
breakTag:
}

// 打印九九乘法表
func practise() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}
