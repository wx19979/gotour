package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum   int
	mutex sync.RWMutex
)

func main() {
	run()
	doOnce()
	race()
	syncMap()
}

// 模拟跑步场景的函数
func run() {
	var wg sync.WaitGroup //创建等待信号量
	//因为要监控110个协程，所以设置计数器为110
	wg.Add(110)
	for i := 0; i < 100; i++ { //逐次的创建协程
		go func() {
			//计数器值减1
			defer wg.Done()
			add(10)
		}()
	}
	for i := 0; i < 10; i++ { //循环操作进行相加操作
		go func() {
			//计数器值减1
			defer wg.Done()
			fmt.Println("和为:", readSum())
		}()
	}
	//一直等待，只要计数器值为0
	wg.Wait()
}

// 单例操作
func doOnce() {
	//创建只执行一次的信号量
	var once sync.Once
	onceBody := func() { //创建一次执行函数
		fmt.Println("Only once")
	}
	done := make(chan bool)   //创建一个通道
	for i := 0; i < 10; i++ { //模拟执行10次
		go func() { //创建协程调用执行函数
			once.Do(onceBody)
			done <- true //修改done的通道值为true
		}()
	}
	for i := 0; i < 10; i++ { //依次输出done通道内容
		<-done
	}
}

// 10个人赛跑,1个裁判发号施令
func race() {
	//创建加锁的信号量
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock()
			cond.Wait() //等待发令枪响
			fmt.Println(num, "号开始跑……")
			cond.L.Unlock()
		}(i)
	}
	//等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast() //发令枪响
	}()
	wg.Wait()
}

// 创建一个并发的map结构体
func syncMap() {
	syncMap := sync.Map{}
	syncMap.Store(1, 1) //放入信息
	syncMap.Store(1, 2)
	fmt.Println(syncMap.Load(1)) //输出第一map的值
}

// 相加操作
func add(i int) {
	mutex.Lock()         //加锁
	defer mutex.Unlock() //所有操作结束开锁
	sum += i             //执行相加操作(临界区)
}

// 读操作
func readSum() int {
	mutex.RLock()         //加锁
	defer mutex.RUnlock() //所有操作结束解锁
	b := sum              //获取sum
	return b              //返回sum
}
