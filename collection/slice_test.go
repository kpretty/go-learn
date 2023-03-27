package collection

import (
	"fmt"
	"testing"
)

func TestSliceDemo1(t *testing.T) {
	sliceDemo1()
}

func TestSliceDemo2(t *testing.T) {
	sliceDemo2()
}

func TestGetSlice(t *testing.T) {
	getSlice()
}
func TestSliceEssence(t *testing.T) {
	sliceEssence()
}

func TestCopySlice(t *testing.T) {
	copySlice()
}

func TestAppendSlice(t *testing.T) {
	appendSlice()
}

func TestDeleteSlice(t *testing.T) {
	deleteSlice()
}

func TestMyDelete(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s1 = myDelete(s1, 3)
	fmt.Println(s1)
}
