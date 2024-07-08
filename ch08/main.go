// 通道的测试代码
package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个通道对象
	ch := make(chan string)
	//执行一个go的协程将信息放入通道之中
	go func() {
		fmt.Println("飞雪无情")
		ch <- "goroutine 完成"
	}()
	//打印结果输出当前是主函数的协程信息
	fmt.Println("我是 main goroutine")
	//将通道ch的信息放到变量v中
	v := <-ch
	//输出具体的信息
	fmt.Println("接收到的chan中的值为:", v)
	//创建一个缓冲区大小是5的一个通道
	cacheCh := make(chan int, 5)
	//放入两个数据
	cacheCh <- 2
	cacheCh <- 3
	//输出该通道的信息
	fmt.Println("cacheCh容量为:", cap(cacheCh), ",元素个数为：", len(cacheCh))
	//创建三个string类型的通道
	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)
	//三个协程分被执行下载文件操作
	go func() {
		firstCh <- downloadFile("firstCh")
	}()

	go func() {
		secondCh <- downloadFile("secondCh")
	}()

	go func() {
		threeCh <- downloadFile("threeCh")
	}()
	//通过select语句实现并发
	select {
	case filePath := <-firstCh: //当为第一路径时输出内容
		fmt.Println(filePath)
	case filePath := <-secondCh: //当为第二路径时输出内容
		fmt.Println(filePath)
	case filePath := <-threeCh: //当为第三路径时输出内容
		fmt.Println(filePath)
	}
}

// 模拟下载环境的函数
func downloadFile(chanName string) string {
	//模拟下载文件,可以自己随机time.Sleep点时间试试
	time.Sleep(time.Second)
	return chanName + ":filePath"
}
