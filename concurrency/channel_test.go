package concurrency

import (
	"testing"
)

func TestInitChannel(t *testing.T) {
	initChannel()
}

func TestUseChannel1(t *testing.T) {
	useChannel1()
}

func TestUseChannel2(t *testing.T) {
	useChannel2()
}

func TestUseChannel3(t *testing.T) {
	useChannel3()
}

func TestSingleChannel(t *testing.T) {
	singleChannel()
}

func TestForCloseChannel(t *testing.T) {
	forCloseChannel()
}
