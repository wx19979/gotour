// 如果对一个变量赋值,这个变量必须有对应分配好的内存这样
// 才可以对这块内存进行操作,完成赋值的目的
// 一个变量必须要通过声明、内存分配才能够赋值,才可以在声明时
// 进行初始化
// 对于值类型的变量即使没有初始化,也可以在之后重新使用该变量
// 但是对于指针类型数据则不同,如果不分配对应的内存地址,那么
// 该指针默认值是nil后面根本无法赋值
// new函数是根据传入参数的类型申请对应类型的内存空间,但是初始
// 化的值是对应类型的零值。new函数只用于分配内存,且把内存清零
// 返回一个指向对应类型零值的指针。
// make函数是map、slice、chan三种类型的工厂函数同时只用于这
// 三种类型变量的创建和初始化,因为这三种类型的结构比较复杂。
package main

import "fmt"

func main() {
	var s string           //创建一个字符串
	fmt.Printf("%p\n", &s) //输出字符串的地址
	s = "张三"               //创建一个字符串
	fmt.Println(s)         //输出字符串具体内容

	var sp *string = nil //创建一个字符串的指针
	sp = new(string)     //申请一个字符串内存空间
	*sp = "飞雪无情"         //设置数据内容
	fmt.Println(*sp)

	pp := NewPerson("飞雪无情", 20)                    //根据构造函数创建一个人的对象
	fmt.Println("name为", pp.name, ",age为", pp.age) //输出结果

	m := map[string]int{"张三": 18} //创建一个map对象
	fmt.Println(m)                //输出信息

}

//创建人对象的函数(工厂函数)
func NewPerson(name string, age int) *person {
	p := new(person) //申请空间
	p.name = name    //设置姓名
	p.age = age      //设置年龄
	return p         //返回对象
}

// 人对象结构体
type person struct {
	name string
	age  int
}
