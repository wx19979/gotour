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
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
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
