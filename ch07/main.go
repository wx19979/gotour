// 错误和异常
package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// 测试这个函数调用是否出错
	i, err := strconv.Atoi("a")
	if err != nil { //如果有问题直接输出错误信息
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	sum, err := add(-1, 2) //调用相加函数
	var cm *commonError
	if cm, ok := err.(*commonError); ok { //错误类型断言
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else { //否则输出正确结果
		fmt.Println(sum)
	}
	if errors.As(err, &cm) { //判断当前错误如果符合该条件输出错误结果
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else { //否则输出正确结果
		fmt.Println(sum)
	}
	//创建一个错误的对象
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e) //包裹嵌套成新的error对象错误信息(error.Unwrap解开嵌套)
	fmt.Println(w)
	//输出错误信息
	fmt.Println(errors.Unwrap(w))
	fmt.Println(errors.Is(w, e))
	//调用defer函数
	moreDefer()
	// go语言中可以通过内置的recover函数恢复panic异常
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	//调用连接数据库函数
	connectMySQL("", "root", "123456")
}

// 自己定义的相加函数(可以返回错误信息以及加工的数据)
func add(a, b int) (int, error) {
	if a < 0 || b < 0 { //如果a或者b为负数返回错误信息的结构体
		return 0, &commonError{
			errorCode: 1,           //错误代码
			errorMsg:  "a或者b不能为负数"} //具体错误信息
	} else { //如果正常情况就返回相加的结果
		return a + b, nil
	}
}

// 模拟连接服务器的函数
func connectMySQL(ip, username, password string) {
	if ip == "" { //如果ip地址是空的状态
		panic("ip不能为空") //抛出一个异常
	}
	//省略其他代码
}

// defer测试函数(defer函数保证文件关闭后一定会被执行,不管自定义的函数出现异常还是报错)
// 一个函数中可以有多个defer语句,多个defer语句顺序按照后进先出的顺序执行
func moreDefer() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Three defer")
	fmt.Println("函数自身代码")
}

// 自定义的错误结构体
type commonError struct {
	errorCode int    //错误码
	errorMsg  string //错误信息
}

// 实现Error这个接口的函数
func (ce *commonError) Error() string {
	return ce.errorMsg
}
