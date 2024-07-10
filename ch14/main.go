// 1,go语言传参都是值传递
// 值传递指的是传递原来数据的一份拷贝而不是原来数据的本身
// 2,指针类型的变量保存的是数据对应内存的地址,所以在函数参
// 数传值的原则下拷贝的值也是内存地址
// 3,值传递的是指针也就是内存地址,通过内存地址可以找到原数据的
// 那块内存,修改它相当于就是修改原数据
// 4,go语言通过make函数或者字面量的包装省了指针的操作,使可以更
// 容易地使用map,其实就是语法糖（在生成map的时候都是调用一个叫
// makemap的函数而里面的参数实际还是指针）注意:这里map可以理解
// 为引用类型,在参数传递时,它还是值传递
// 5,go语言没有引用类型,但是可以把map,chan称之为引用类型
// 除了map、chan外,go语言中的函数、接口、slice切片都可以称为引用类型
// 指针类型也可以理解为时一种引用类型
// 6,go语言中,定义通过声明或通过make和new函数属于显式声明并初始化,
// 如果声明变量没有显式声明初始化那该变量默认值就是对应类型的零值
// go语言中只用值传递,且传递的是原始数据的拷贝,如果拷贝内容是值类型,
// 那么函数无法修改原始数据,如果拷贝的是指针类型则函数可以修改原始数据
// 并且go语言严格来说只分为值类型和指针类型两种类型
package main

import (
	"fmt"
)

func main() {
	//创建一个人的对象
	p := person{name: "张三", age: 18}
	fmt.Printf("main函数：p的内存地址为%p\n", &p)
	//调用修改函数
	modifyPerson(&p)
	fmt.Println("person name:", p.name, ",age:", p.age)
	//创建一个map的通道
	m := make(map[string]int)
	m["飞雪无情"] = 18
	fmt.Println("飞雪无情的年龄为", m["飞雪无情"])
	fmt.Printf("main函数：m的内存地址为%p\n", m)
	modifyMap(m) //修改map
	fmt.Println("飞雪无情的年龄为", m["飞雪无情"])

	//类型零值测试
	var s string
	var i int
	var b bool
	var f float64
	var st struct{}
	var mi map[string]int
	var sl []string
	var ia interface{}
	var fn func()
	var ch chan string
	fmt.Println("string的零值为", s)
	fmt.Println("int的零值为", i)
	fmt.Println("bool的零值为", b)
	fmt.Println("float64的零值为", f)
	fmt.Println("struct的零值为", st)
	fmt.Println("map的零值为", mi)
	fmt.Println("slice的零值为", sl)
	fmt.Println("interface的零值为", ia)
	fmt.Println("func的零值为", fn)
	fmt.Println("chan的零值为", ch)
}

func modifyPerson(p *person) {
	fmt.Printf("modifyPerson函数：p的内存地址为%p\n", p)
	p.name = "李四"
	p.age = 20
}

func modifyMap(p map[string]int) {
	fmt.Printf("modifyMap函数：p的内存地址为%p\n", p)
	p["飞雪无情"] = 20
}

// 人类的结构体
type person struct {
	name string //人名
	age  int    //年龄
}
