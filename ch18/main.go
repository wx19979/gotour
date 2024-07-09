package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	ss := []string{"飞雪无情", "张三"}                     //创建一个字符串切片
	fmt.Println("切片ss长度为", len(ss), ",容量为", cap(ss)) //输出当前切片的信息
	ss = append(ss, "李四", "王五")                      //再在切片后面加上信息
	fmt.Println("切片ss长度为", len(ss), ",容量为", cap(ss)) //输出修改过后切片的信息
	fmt.Println(ss)

	a1 := [2]string{"飞雪无情", "张三"}      //创建一个切片对象
	fmt.Printf("函数main数组指针：%p\n", &a1) //输出主函数指针
	arrayF(a1)                         //调用函数

	s1 := a1[0:1]                                                 //截取切片
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data) //输出最终反射获取到的结果
	sliceF(s1)                                                    //调用函数

	s2 := a1[:]                                                   //获取完整的切片
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data) //输出获取的数据

	sh1 := (*slice)(unsafe.Pointer(&s1))    //将当前的内容强制转换为切片对象
	fmt.Println(sh1.Data, sh1.Len, sh1.Cap) //输出切片的信息

	s := "飞雪无情"                                                                 //创建一个字符串的对象
	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data) //输出地址的信息
	b := []byte(s)                                                              //创建byte切片的对象
	fmt.Printf("b的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)  //输出地址的信息
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))                            //强制转换为切片指针类型
	sh.Cap = sh.Len                                                             //修改当前容量为其对应对象的长度
	b1 := *(*[]byte)(unsafe.Pointer(sh))                                        //将指针强制转化为byte数组的指针
	fmt.Printf("b1的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b1)).Data)
	s3 := string(b)
	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)
	s4 := *(*string)(unsafe.Pointer(&b))
	fmt.Printf("s4的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s4)).Data)
	fmt.Println(s, string(b), string(b1), s3)
}

// 根据格式输出指针信息的函数
func arrayF(a [2]string) {
	fmt.Printf("函数arrayF数组指针：%p\n", &a)
}

// 根据格式输出切片信息的函数
func sliceF(s []string) {
	fmt.Printf("函数sliceF Data：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
}

// 自己定义一个切片的结构体
type slice struct {
	Data uintptr //数据字段
	Len  int     //长度字段
	Cap  int     //容量字段
}
