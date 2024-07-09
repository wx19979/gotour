// sync.Mutex 互斥锁,sync.RWMutex读写锁,sync.WaitGroup属于等待组，用于监听所有协程操作的信号量,协调多个协程共同去做一件事情
// sync.Once让代码只执行一次,哪怕是高并发情况下，比如创建一个单例。
// sync.Cond可以用于发号施令关键点在于协程开始的时候是等待的,从字面意思理解是条件变量,它具有阻塞协程和唤醒协程的功能,所以可以在满足一定条件的情况下唤醒协程
// sync.Cond共有三个方法:(1)Wait,阻塞当前协程(2)Signal,唤醒一个等待时间最长的协程(3)Broadcast,唤醒所有等待的协程
// sync.Map并发安全的映射结构,有如下5个方法:
// (1)Store:存储一对key-value值 (2)Load:根据key获取对应的value值,并可以判断key是否存在(3)LoadOrStore:如果key对应value值存在,则返回value如果不存在
// 存储相应的value(4)Delete:删除一个key-value对(5)Range:循环迭代sync.Map,效果与for range一样
// 临界区:访问共享资源的程序片段,而这些共享资源又无法同时被多个协程访问的特性
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
	//因为要监控110个协程，所以设置计数器为110,要跟踪多少协程就设置多少
	wg.Add(110)
	for i := 0; i < 100; i++ { //逐次的创建协程
		go func() {
			//计数器值减1,告诉等待组协程减一
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
	//一直等待，只有计数器值为0,等待这个动作才能结束
	wg.Wait()
}

// 单例操作
func doOnce() {
	//创建只执行一次的信号量
	var once sync.Once
	onceBody := func() { //创建一次执行函数
		fmt.Println("Only once")
	} //函数不管开了几个协程只能被执行一次
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
	//创建一个cond指针用于阻塞和唤醒协程
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ { //先将所有协程阻塞起来等待发号施令
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
		cond.Broadcast() //发令枪响,通知所有协程开始跑
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
