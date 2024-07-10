// 指针的使用环境:
// 1,不要对map,slice,channel这类引用类型使用指针
// 2,如果需要修改方法接收者内部的数据或者状态时,需要使用指针
// 3,如果需要修改参数的值或者内部数据时,也需要使用指针类型的参数
// 4,如果是比较大的结构体这个时候可以考虑用指针
// 5,像int,bool这样的小数据类型没必要用指针
// 6,如果需要并发安全,尽可能的不使用指针,使用指针一定要保证并发安全
// 7,指针最好不要嵌套,也就是不要使用一个指针的指针
// 对是否使用指针作为接收者,有如下几点:
// 1,如果接收者是map,slice,channel这类引用类型,不使用指针
// 2,如果需要修改接收者,那么需要使用指针
// 3,如果接收者是比较大的类型,可以考虑使用指针
// 指针的两大好处
// 1,可以修改指向数据的值
// 2,在变量赋值,参数传值的时候节省内存
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	name := "飞雪无情"
	nameP := &name //取地址
	fmt.Println("name变量的值为:", name)
	fmt.Println("name变量的内存地址为:", nameP, &nameP)

	nameV := *nameP
	fmt.Println("nameP指针指向的值为:", nameV)

	*nameP = "公众号:飞雪无情" //修改指针指向的值
	fmt.Println("nameP指针指向的值为:", *nameP)
	fmt.Println("name变量的值为:", name)

	age := 18
	modifyAge(&age)
	fmt.Println("age的值为:", age)

	var w io.Writer = os.Stdout
	wp := &w
	fmt.Println(wp)
}

// 传递指针修改的函数(只有指针参数才能修改值)
func modifyAge(age *int) {
	*age = 20
}
