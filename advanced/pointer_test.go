package advanced

import (
	"fmt"
	"testing"
)

func TestPointerNature(t *testing.T) {
	pointerNature()
}

func TestGetPointerValue(t *testing.T) {
	a := 1
	fmt.Println(getPointerValue(&a))
}

func TestInitError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("initError panicked with error: %v\n", r)
		}
	}()
	initError()
}

func TestInitSuccess(t *testing.T) {
	initSuccess()
}
