package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	result := make(chan string)
	go func() {
		//模拟网络访问
		time.Sleep(8 * time.Second)
		result <- "服务端结果"
	}()
	//select timeout模式
	select {
	case v := <-result:
		fmt.Println(v)
	case <-time.After(5 * time.Second):
		fmt.Println("网络访问超时了")
	}
	// Pipline模式即流水线模式
	coms := buy(100) //采购100套配件
	//三班人同时组装100部手机
	phones1 := build(coms)
	phones2 := build(coms)
	phones3 := build(coms)
	//汇聚三个channel成一个,扇入扇出模式
	phones := merge(phones1, phones2, phones3)
	packs := pack(phones) //打包它们以便售卖

	//输出测试，看看效果
	for p := range packs {
		fmt.Println(p)
	}
	//Futures模式
	vegetablesCh := washVegetables() //洗菜
	waterCh := boilWater()           //烧水

	fmt.Println("已经安排洗菜和烧水了，我先眯一会")
	time.Sleep(2 * time.Second)
	fmt.Println("要做火锅了，看看菜和水好了吗")
	vegetables := <-vegetablesCh
	water := <-waterCh
	fmt.Println("准备好了,可以做火锅了:", vegetables, water)

}

// 工序1采购
func buy(n int) <-chan string {
	out := make(chan string) //创建输出通道
	go func() {              //协程执行采购操作
		defer close(out)          //执行之后关闭通道
		for i := 1; i <= n; i++ { //将配件放入通道中
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out //返回输出的通道
}

// 工序2组装
func build(in <-chan string) <-chan string {
	out := make(chan string) //创建一个输出的通道
	go func() {              //协程进行组装操作
		defer close(out)    //所有执行完之后关闭通道
		for c := range in { //执行组装操作
			out <- "组装(" + c + ")"
		}
	}()
	return out //返回输出通道
}

// 工序3打包
func pack(in <-chan string) <-chan string {
	//创建打包的通道
	out := make(chan string)
	go func() { //协程执行打包操作
		defer close(out)    //所有执行完之后关闭通道
		for c := range in { //循环执行打包操作
			out <- "打包(" + c + ")" //将打包信息放入通道中
		}
	}()
	return out //返回输出的打包通道
}

// 扇入函数（组件）,把多个chanel中的数据发送到一个channel中
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup    //创建一个同步等待信号组
	out := make(chan string) //创建一个输出的通道

	//把一个channel中的数据发送到out中
	p := func(in <-chan string) {
		defer wg.Done()     //执行所有之后关闭通道
		for c := range in { //循环输出操作
			out <- c
		}
	}

	wg.Add(len(ins))

	//扇入，需要启动多个goroutine用于处于多个channel中的数据
	for _, cs := range ins {
		go p(cs)
	}

	//等待所有输入的数据ins处理完，再关闭输出out
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// 洗菜
func washVegetables() <-chan string {
	//创建一个信息通道
	vegetables := make(chan string)
	go func() { //协程进行洗菜操作
		time.Sleep(5 * time.Second) //每隔5秒执行操作
		vegetables <- "洗好的菜"        //将信息放入通道中
	}()
	return vegetables //返回信息
}

// 烧水
func boilWater() <-chan string {
	water := make(chan string) //创建一个烧水信息的通道
	go func() {                //协程进行烧水操作
		time.Sleep(5 * time.Second) //每隔5秒进行操作
		water <- "烧开的水"             //将信息放入通道中
	}()
	return water //返回信息
}
