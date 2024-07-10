// Context用于跟踪每一个协程的状态,能更好的控制协程
// Context是一个接口,具有手动、定时、超时发出取消信号、传值信号等功能,主要用于控制多个协程之间的协作,尤其是取消操作
// Context共有4个函数(1)Deadline获取设置的截止时间(2)Done返回一个只读的channel,类型为struct{}(3)Err返回取消错误原因
// (4)Value获取Context上绑定的值,是一个键值对
// Context树的几个函数:(1)Background:用于创建Context根节点对象(2)WithCancel(parent Context):生成一个可取消的Context
// (3)WithDeadline(parent Context,d time.Time):生成一个可定时取消的Context,参数d为定时取消的具体时间
// (4)WithTimeout(parent Context,timeout time.Duration):生成一个可超时取消的Context,参数timeout用于设置多久后取消
// (5)WithValue(parent Context,key,val interface{}):生成一个可携带key-value键值对的Context
// Context使用原则:
// (1)Context不要放到结构体中,要以参数的方式传递
// (2)Context作为函数参数的时候要放在第一位,也就是第一个参数
// (3)要使用Context.Background函数生成根节点的Context也就是最顶层的Context
// (4)Context传值要传递必须的值而且尽可能的少不要什么都传
// (5)Context多协程安全,可以在多个协程中放心使用
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//创建用于进行并发等待对象
	var wg sync.WaitGroup
	wg.Add(4) //初始化等待组对象
	//创建一个ctx对象 context.Background()用于创建一个空的context对象作为整个context的根节点
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
	//通过context进行传值,先将key-value对存储到context中
	valCtx := context.WithValue(ctx, "userId", 2)
	go func() { //并发调用获取用户信息函数
		defer wg.Done()
		getUser(valCtx)
	}()

	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stop()                      //发停止指令,该函数属于context的内部函数用于发送结束的这一信息
	wg.Wait()
}

// 监控狗函数
func watchDog(ctx context.Context, name string) {
	//开启for select循环，一直后台监控
	for {
		select {
		case <-ctx.Done(): //当信号量为结束信号量(对应就是context调用了stop函数)
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
		default: //取出context中的key所对应的value值
			userId := ctx.Value("userId")
			fmt.Println("【获取用户】", "用户ID为：", userId) //输出信息
			time.Sleep(1 * time.Second)             //等待1秒钟
		}
	}
}
