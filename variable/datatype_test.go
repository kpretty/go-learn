package variable

import (
	"fmt"
	"testing"
)

func TestIntType(t *testing.T) {
	fmt.Printf("i1 type [%T] value [%d]\n", i1, i1)
	fmt.Printf("i2 type [%T] value [%d]\n", i2, i2)
	fmt.Printf("i3 type [%T] value [%d]\n", i3, i3)
	fmt.Printf("i4 type [%T] value [%d]\n", i4, i4)
	fmt.Printf("i5 type [%T] value [%d]\n", i5, i5)
	fmt.Printf("i6 type [%T] value [%d]\n", i6, i6)
	fmt.Printf("i7 type [%T] value [%d]\n", i7, i7)
	fmt.Printf("i8 type [%T] value [%d]\n", i8, i8)
	fmt.Printf("i9 type [%T] value [%d]\n", i9, i9)
	fmt.Printf("i10 type [%T] value [%d]\n", i10, i10)
	fmt.Printf("i11 type [%T] value [%d]\n", i11, i11)
}

func TestNumberLiteral(t *testing.T) {
	fmt.Printf("v1 binary [%b] value [%d]\n", v1, v1)
	fmt.Printf("v2 binary [%b] value [%d]\n", v2, v2)
	fmt.Printf("v3 binary [%b] value [%d]\n", v3, v3)
	fmt.Printf("v4 binary [%b] value [%d]\n", v4, v4)
}

func TestFloat(t *testing.T) {
	fmt.Printf("f1 type [%T] %f\n", f1, f1)
	fmt.Printf("f2 type [%T] %.2f\n", f2, f2)
	fmt.Printf("f1 type [%T] %f\n", f3, f3)
	fmt.Printf("f1 type [%T] %f\n", f4, f4)
}

func TestComplex(t *testing.T) {
	fmt.Printf("f1 type [%T] %f\n", c1, c1)
	fmt.Printf("f1 type [%T] %f\n", c2, c2)
	fmt.Printf("f1 type [%T] %f\n", c1*c2, c1*c2)
}

func TestChar(t *testing.T) {
	fmt.Printf("char1 type [%T] value %c\n", char1, char1)
	fmt.Printf("char1 type [%T] value %c\n", char2, char2)
}

func TestString(t *testing.T) {
	fmt.Printf(" s1 type [%T] value %s\n", s1, s1)
	verifyStringIsByteArray()
}

func TestEditString(t *testing.T) {
	editString()
}

func TestTypeTransform(t *testing.T) {
	typeTransform()
}

func TestStringConnection(t *testing.T) {
	stringsConnection()
}
