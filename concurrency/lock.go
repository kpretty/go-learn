package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var x int

func add() {
	for i := 0; i < 5000; i++ {
		x += 1
	}
	wg.Done()
}

var lock sync.Mutex

func addSolve() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

func MutexProblem() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x) // 解决是随机的

}

// 互斥锁，当一个goroutine获取到一把互斥锁，后续其他goroutine尝试获取这把锁时将发生阻塞，直到锁被释放
// go 获取锁的机制默认是非公平的，即所有goroutine随机获取锁

func MutexSolve() {
	wg.Add(2)
	go addSolve()
	go addSolve()
	wg.Wait()
	fmt.Println(x)
}

// 读写锁
// 读共享、写互斥：
//		当一个goroutine获取到写锁时，其他goroutine将无法获取到其他锁（读锁写锁）发生阻塞，直到写锁被释放
// 		当一个goroutine获取到读锁时，其他goroutine可以继续获取读锁实现数据的读取，但当一个goroutine尝试获取写锁时会被阻塞，直到所有读锁被释放再重新尝试
// 读写锁获取锁机制默认也是非公平的，即所有goroutine随机获取锁
// 场景：读多写少
// 优点：提高吞吐量，因为读操作的时间远远小于写操作
// 缺点：写操作可能会被无限延迟
var rwLock sync.RWMutex
var rwX int

func readWithRWLock(wg *sync.WaitGroup) {
	// 获取读锁
	rwLock.RLock()
	defer rwLock.RUnlock()
	// 执行读操作，假设读操作耗时一毫秒
	time.Sleep(time.Millisecond)
	wg.Done()
}

func WriteWithRWLock(wg *sync.WaitGroup) {
	// 获取写锁
	rwLock.Lock()
	defer rwLock.Unlock()
	// 执行写操作，假设写操作耗时十毫秒
	time.Sleep(10 * time.Millisecond)
	rwX += 1
	wg.Done()
}

// 对照组，互斥锁
func readWithLock(wg *sync.WaitGroup) {
	// 获取互斥锁
	lock.Lock()
	defer lock.Unlock()
	// 执行读操作，假设读操作耗时一毫秒
	time.Sleep(time.Millisecond)
	wg.Done()
}

func WriteWithLock(wg *sync.WaitGroup) {
	// 获取写锁
	lock.Lock()
	defer lock.Unlock()
	// 执行写操作，假设写操作耗时十毫秒
	time.Sleep(10 * time.Millisecond)
	rwX += 1
	wg.Done()
}

// Go 用于执行对照组实验
// rf wf 读写函数
// rc wc 读写次数
func Go(rf, wf func(wg *sync.WaitGroup), rc, wc int) {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf(&wg)
	}
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf(&wg)
	}
	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x=%v, cost:%v\n", rwX, cost)
}

// sync.Once 永远只会执行一次
// 单例模式

var config map[string]string
var loadConfig sync.Once

func initConfig() {
	config = map[string]string{
		"1": "1",
		"2": "2",
	}
	fmt.Println("加载完成")
}

// getConfig 高并发情况下会存在多次加载、数据异常情况
func getConfigWithoutSingle(key string) string {
	if config == nil {
		initConfig()
	}
	return config[key]
}

func getConfigWithSingle(key string) string {
	loadConfig.Do(initConfig)
	return config[key]
}

func GoSingle1() {
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println(getConfigWithoutSingle("1"))
		}()
	}
}

func GoSingle2() {
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println(getConfigWithSingle("1"))
		}()
	}
}
