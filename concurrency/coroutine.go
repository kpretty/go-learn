package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// 无畏并发

func step1() {
	time.Sleep(5 * time.Second)
	fmt.Println("第一步完成")
}

func step2() {
	time.Sleep(5 * time.Second)
	fmt.Println("第二步完成")
}

// Step1 串行执行，需要耗时 10s
func Step1() {
	fmt.Println("任务开始")
	step1()
	step2()
	fmt.Println("任务结束")
}

// Step2 启用 goroutine 但不会打印step1 2 因为由 main goroutine 创建的两个 goroutine 还没执行完 main 就结束了
// 当 main goroutine 结束时凡是由其创建的 goroutine 将自动结束，可以在 main 中进行 Sleep 但显然是不合理的
func Step2() {
	fmt.Println("任务开始")
	go step1()
	go step2()
	fmt.Println("任务结束")
}

// 使用 sync 的并发原语，类似 java 的CountDownLatch
// 声明全局等待组
var wg sync.WaitGroup

func step1_() {
	time.Sleep(5 * time.Second)
	fmt.Println("第一步完成")
	wg.Done()
}

func step2_() {
	time.Sleep(5 * time.Second)
	fmt.Println("第二步完成")
	wg.Done()
}

// Step3 并行，耗时 5s
func Step3() {
	wg.Add(2)
	fmt.Println("任务开始")
	go step1_()
	go step2_()
	fmt.Println("任务结束")
	// 等待所有 goroutine 完成
	wg.Wait()
}

func multiStep(step int) {
	// 无论后续会发生什么都会 Done
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("第 %d 步完成\n", step)
}

// Step4 并行，几乎耗时 1s
func Step4() {
	start := time.Now().UnixMilli()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go multiStep(i)
	}
	wg.Wait()
	end := time.Now().UnixMilli()
	fmt.Println("耗时：", end-start)
	time.Now().Format("")
}
