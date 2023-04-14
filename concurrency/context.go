package concurrency

import (
	"context"
	"fmt"
	"time"
)

// 并发编程重要思想：如果你无法掌控你创建的线程，那就不要创建
// 在之前的知识点中，借助关键字 go 可以随心所欲的创建 goroutine，但如何去治理这些 goroutine 是个难题
// 目前掌握的知识多个goroutine之间通信只能依赖 channel，如果一个 goroutine 创建若干个子 goroutine 如何去管理
// 似乎只能为每个子 goroutine 创建对应的信号通道 chan struct{} 当想要停止 goroutine 向通道发送数据即可
// 但这样需要维护很多 channel 且跨包不容易管理
// 使用 context 来优雅管理 goroutine 生命周期

// 简单使用

func worker(ctx context.Context, name string) {
LOOP:
	for true {
		select {
		case <-ctx.Done():
			fmt.Printf("%v要退出了，因为：%v\n", name, ctx.Err())
			break LOOP
		default:
			fmt.Printf("%v 在工作\n", name)
			time.Sleep(time.Second)
		}
	}
}

func Master() {
	// 创建 context 用来控制 worker
	ctx, cancelFunc := context.WithCancel(context.Background())
	// 只要调用 cancelFunc 凡是由当前 goroutine 创建的 goroutine 都会被终止
	go worker(ctx, "work1")
	go worker(ctx, "work2")
	go worker(ctx, "work3")

	time.Sleep(3 * time.Second)
	fmt.Println("worker们，你们下班了")
	// master 表示你们可以下班了，停掉所有我创建的 goroutine
	cancelFunc()
	for i := 0; i < 3; i++ {
		fmt.Println("认真工作的master")
		time.Sleep(time.Second)
	}
	fmt.Println("master我也要下班了")
}
