package collection

import (
	"fmt"
	"testing"
)

func TestArraySwap(t *testing.T) {
	arraySwap()
}

func TestDefineArray(t *testing.T) {
	defineArray()
}

func TestArrayIsValuePass(t *testing.T) {
	src1 := [...]int{1, 2, 3, 4, 5}
	src2 := src1[0 : len(src1)-1]
	arrayIsValuePass(src1)
	fmt.Println(src1)
	sliceIsValuePass(src2)
	fmt.Println(src2)
}

// 两数之和
func TestTwoNumSum(t *testing.T) {
	fmt.Println(twoNumSum([]int{11, 22, 13, 24, 15, 26, 17}, 46))
}
