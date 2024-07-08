package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//创建勇于进行并发等待对象
	var wg sync.WaitGroup
	wg.Add(4) //初始化等待组对象
	//创建一个ctx对象
	ctx, stop := context.WithCancel(context.Background())
	// 开启三个协程调用监控狗的函数
	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗1】")
	}()

	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗2】")
	}()

	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗3】")
	}()
	//上下文调用当前的值
	valCtx := context.WithValue(ctx, "userId", 2)
	go func() { //并发调用获取用户信息函数
		defer wg.Done()
		getUser(valCtx)
	}()

	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stop()                      //发停止指令
	wg.Wait()
}

// 监控狗函数
func watchDog(ctx context.Context, name string) {
	//开启for select循环，一直后台监控
	for {
		select {
		case <-ctx.Done(): //当信号量为结束信号量
			fmt.Println(name, "停止指令已收到，马上停止") //输出信息
			return                            //直接返回
		default: //正常情况下输出正在监控信息
			fmt.Println(name, "正在监控……")
		}
		time.Sleep(1 * time.Second) //等待1秒钟
	}
}

// 获取用户信息的函数
func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //当获取到结束信号量
			fmt.Println("【获取用户】", "协程退出") //提示当前协程结束
			return
		default: //正常运行状态
			userId := ctx.Value("userId")           //获取到用户的ID
			fmt.Println("【获取用户】", "用户ID为：", userId) //输出信息
			time.Sleep(1 * time.Second)             //等待1秒钟
		}
	}
}
