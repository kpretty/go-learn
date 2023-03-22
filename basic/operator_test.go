package basic

import (
	"fmt"
	"testing"
)

func TestFindOnceNumber(t *testing.T) {
	fmt.Println(findOnceNumber([]int{1, 1, 2, 2, 3, 4, 4}))
}
