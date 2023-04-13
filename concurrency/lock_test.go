package concurrency

import (
	"runtime"
	"testing"
)

func TestMutexProblem(t *testing.T) {
	MutexProblem()
}

func TestMutexSolve(t *testing.T) {
	MutexSolve()
}

func TestGo(t *testing.T) {
	rwX = 0
	Go(readWithRWLock, WriteWithRWLock, 1000, 10) // x=10, cost:111.6855ms
	rwX = 0
	Go(readWithLock, WriteWithLock, 1000, 10) // x=10, cost:1.264073292s
	// 效率上 十倍差距，随着读次数逐渐增大互斥锁的效率将越来越低，读写锁效率几乎不变
}

func TestGoSingle1(t *testing.T) {
	runtime.GOMAXPROCS(8)
	GoSingle1()
}

func TestGoSingle2(t *testing.T) {
	runtime.GOMAXPROCS(8)
	GoSingle2()
}
