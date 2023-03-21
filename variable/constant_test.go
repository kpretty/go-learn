package variable

import (
	"fmt"
	"testing"
)

func TestConstDeclaration(t *testing.T) {
	fmt.Println("π = ", pi)
	fmt.Println("e = ", e)
	fmt.Println("c10 = ", c10)
}

func TestIota(t *testing.T) {
	fmt.Println("n1 = ", n1)
	fmt.Println("n2 = ", n2)
	fmt.Println("n3 = ", n3)
	fmt.Println("n4 = ", n4)
	fmt.Println("n5 = ", n5)
}

func TestIotaCase(t *testing.T) {
	fmt.Println("KB = ", KB)
	fmt.Println("MB = ", MB)
	fmt.Println("GB = ", GB)
	fmt.Println("TB = ", TB)
	fmt.Println("PB = ", PB)
}
