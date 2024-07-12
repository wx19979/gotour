// 一个数组由数组的大小和数组内的元素类型构成,一旦一个数组被声明，它的大小和内部元素的类型就不能改变
// Go语言中,函数间的传递是值的传递,数组作为参数在各函数间被传递时同内容会被一遍一遍复制,这就造成大量内存浪费
// 为了解决限制Go语言创造了slice切片,切片是对数组的抽象和封装,它的底层是一个数组存储所有的元素,可把切片理解为动态数组
// 通过内置的append方法,可向一个切片追加任意多个元素可解决数组第一个限制
// append自动扩容原理是新建一个底层数组,把原来切片内的元素拷贝到新数组中,再返回一个指向新数组的切片
// 数组和切片的取值和赋值操作要更高效因为它们是连续的内存操作通过索引可快速找到元素存储的地址
// Go语言通过先分配一个内存再复制内容的方式实现string和[]byte之间的强制转换
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
	//测试用unsafe.Pointer通过零拷贝进行类型互转
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
