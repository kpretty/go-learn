package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// 单纯的启动 goroutine 意义不是很大，需要多个 goroutine 通讯协同运作
// go 的理念是：通过通讯共享内存，而不是通过共享内存实现通讯

// channel 作为 goroutine 的连接，是一个特殊引用类型，遵循 FIFO 规则，channel只能接收一种类型
// var 变量名 chan 元素类型

func initChannel() {
	var c1 chan int
	var c2 chan bool
	var c3 chan []int
	// channel 的零值是 nil
	fmt.Println(c1, c2, c3)
	// channel 的操作有三种：接收、发送、关闭
	// 对 nil 做接收、发送会发生阻塞，关闭 nil 通道会发生 panic
	// 强制运行会发生死锁
	// c1 <- 1
	// i := <-c1
	// fmt.Println(i)
	// close(c1)
}

func useChannel1() {
	// 使用 make 初始化 channel
	// 无缓冲的channel，又称同步 channel，其接收和发送都是同步阻塞的
	c1 := make(chan int)
	wg.Add(2)
	go func() {
		fmt.Println("准备向通道发送一个数据")
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		c1 <- rand.Int()
		fmt.Println("发送完了")
		wg.Done()
	}()

	go func() {
		fmt.Println("准备向通道接收一个数据")
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		e := <-c1
		fmt.Println("接收到数据是：", e)
		wg.Done()
	}()

	wg.Wait()
	close(c1)
}

func useChannel2() {
	// 使用 make 初始化 channel
	// 有缓冲区的 channel，当缓冲区未满时，发送完直接结束无需阻塞，接收数据时当没有数据时会阻塞
	// 当缓冲区满时则退化成同步通道，直到数据被消费
	c1 := make(chan int, 1)
	wg.Add(2)
	go func() {
		fmt.Println("准备向通道发送一个数据")
		// time.Sleep(5 * time.Second)
		c1 <- rand.Int()
		fmt.Println("发送完了")
		wg.Done()
	}()

	go func() {
		fmt.Println("准备向通道接收一个数据")
		time.Sleep(5 * time.Second)
		e := <-c1
		fmt.Println("接收到数据是：", e)
		wg.Done()
	}()

	wg.Wait()
	close(c1)
}

// a 向 c1 发送 0-100
// b 接受 c1 数据做平方后将数据写到 c2
// main 接收 c2 求和输出
func useChannel3() {
	c1 := make(chan int, 100)
	c2 := make(chan int, 100)
	// a
	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Println("a -> c1 : ", i)
			c1 <- i
		}
		close(c1)
	}()
	go func() {
		// 对 channel 使用 for range 会自动判断 channel 是否被 close，否则会阻塞
		for i := range c1 {
			fmt.Println("c1 calc : ", i)
			c2 <- i * i
		}
		close(c2)
	}()
	var result int
	// 使用多值方式手动 break
	for true {
		value, ok := <-c2
		if !ok {
			break
		}
		result += value
	}
	fmt.Println("main goroutine result : ", result)
}

func singleChannel() {
	// 单向 channel 表示这个 channel 只能接收数据或者发送数据
	// 尝试对单向 channel 做操作不允许的操作时会panic
	// c1 := make(<-chan int, 10) // 只能接收数据
	c2 := make(chan<- int, 10) // 只能发送数据
	// c1 <- 1 Invalid operation: c1 <- 1 (send to the receive-only type <-chan int)
	// <- c2 // Invalid operation: <- c2 (receive from the send-only type chan<- int)
	// close(c1) // 只读通道无法被 close，从业务角度来说，只读通道对于当前操作者是被限制的无法越权关闭
	close(c2)
	// 通常我们可以将一个双向通道作为只读通道的入参进行传递，保证方法内部无法修改通道数据
}

func forCloseChannel() {
	// 尝试对一个已经关闭的 channel 进行操作
	c1 := make(chan int, 10)
	for i := 0; i < 5; i++ {
		c1 <- i
	}
	// 关闭
	close(c1)
	// 对关闭的channel发送数据
	// c1 <- 999 // panic: send on closed channel
	// 对关闭的channel接收数据
	for i := 0; i < 10; i++ {
		// 先消费完通道所有值
		// 然后返回通道零值，因此推荐使用 for range 遍历或多值接收
		fmt.Println(<-c1)
	}
	// 对关闭的channel再次关闭
	// close(c1) // panic: close of closed channel
}
