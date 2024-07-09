// 接口是和调用方法的一种约定,它是一个高度抽象的类型,不用和具体实现细节绑定在一起
package main

import (
	"fmt"
)

func main() {
	//创建一个人的对象
	p := person{
		age:  30,
		name: "飞雪无情",
		address: address{
			province: "北京",
			city:     "北京",
		},
	}
	// 输出此人各种信息
	fmt.Println(p.name, p.age)
	fmt.Println(p.province)
	printString(p.address)
	printString(&p)
	//调用构造这个人的函数
	p1 := NewPerson("张三")
	fmt.Println(p1)
	//创建一个stringer的对象
	var s fmt.Stringer
	s = p1
	p2 := s.(*person) //类型断言,如果是一个person指针直接转换为指针类型,不会异常,否则抛出异常
	fmt.Println(p2)
	a, ok := s.(address) //类型断言的多值返回情况
	if ok {              //这个属于是断言成功
		fmt.Println(a)
	} else { //否则断言失败
		fmt.Println("s不是一个address")
	}

	add := address{province: "北京", city: "北京"} //创建一个地址的对象
	printString(add)                           //输出结果
	printString(&add)
}

// 定义一个人的结构体
type person struct {
	name string //名字的字段
	age  uint   //年龄的字段

	address //套用地址的结构体
}

// 用于通过名字构建人对象的函数(工厂函数)
func NewPerson(name string) *person {
	return &person{name: name}
}

// 将人这个对象的名字和年龄输出的函数(让person类型实现了string接口中的String方法)
func (p *person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

// 实现WalkRun这个借口的walk函数
func (p *person) Walk() {
	fmt.Printf("%s能走\n", p.name)
}

// 实现WalkRun这个借口的Run函数
func (p *person) Run() {
	fmt.Printf("%s能跑\n", p.name)
}

// 输出字符串的函数
func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

// 地址结构体
type address struct {
	province string //省字段
	city     string //城市字段
}

// 输出信息的函数(让address类型实现String接口下的String方法)
func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

// 定义一个WalkRun的接口
type WalkRun interface {
	Walk() //定义内部的走的函数
	Run()  //定义内部的跑的函数
}
