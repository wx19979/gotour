// unsafe包,可绕过Go的安全机制,直接对内存进行读写
// unsafe.Pointer:一种特殊意义的指针,可以表示任意类型的地址
// 类似C语言里面的void*指针,全能型的,但是该指针不能进行指针运算
// uintptr:该类型可以对指针进行偏移运算,这样就可以访问特定内存,
// 达到对特定内存进行读写的目的,这是真正内存级别的操作
// 共有*T,unsafe.Pointer,uintptr三种类型的指针
// 指针转换规则:任何类型*T都可转换为unsafe.Pointer,unsafe.Pointer
// 同样也可以转换为任何类型*T,uintptr可转为unsafe.Pointer,同样
// unsafe.Pointer也可转为uintptr类型unsafe.Pointer是一个桥梁
// Sizeof函数可返回一个类型所占用的内存大小,大小与类型有关,和类型
// 对应变量存储内容大小无关

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10  //int类型
	ip := &i //取地址
	//将地址转换为float64类型
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3 //将该指针对应的数据乘三
	fmt.Println(i)

	p := new(person)
	//Name是person的第一个字段不用偏移，即可通过指针修改
	pName := (*string)(unsafe.Pointer(p))
	*pName = "飞雪无情"
	//Age并不是person的第一个字段，所以需要进行偏移，这样才能正确定位到Age字段这块内存，才可以正确的修改
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age))) //是通过指针偏移找到对应的值的操作
	*pAge = 20

	fmt.Println(*p)
	//分别输出测试的值
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int32(10000000)))
	fmt.Println(unsafe.Sizeof(int64(10000000000000)))
	fmt.Println(unsafe.Sizeof(int(10000000000000000)))
	fmt.Println(unsafe.Sizeof(string("飞雪无情")))
	fmt.Println(unsafe.Sizeof([]string{"飞雪u无情", "张三"}))

}

// 人这一个结构体
type person struct {
	Name string //名字
	Age  int    //年龄
}
