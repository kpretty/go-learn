package advanced

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	r := sum(1, 2)
	fmt.Println(r)
	// 函数返回值也可以不接收
	sum(4, 4)
}

func TestRangeAdd(t *testing.T) {
	fmt.Println(rangeAdd(1, 2, 3, 4, 5))
}

func TestCalc(t *testing.T) {
	fmt.Println(calc1(1, 2))
	fmt.Println(calc2(1, 2))
}

func TestFuncType(t *testing.T) {
	funcType()
}

func TestFuncPointer(t *testing.T) {
	funcPointer()
}

func TestCustomDisplay(t *testing.T) {
	customDisplay(1, func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
}

func TestDeferDemo(t *testing.T) {
	//deferDemo(true)
	deferDemo(false)
}

func TestPractice(t *testing.T) {
	practice()
}
