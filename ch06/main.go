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
	p2 := s.(*person) //将该对象强制转换为字符串
	fmt.Println(p2)
	a, ok := s.(address) //获取s中的地址这一项
	if ok {
		fmt.Println(a)
	} else {
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

// 用于通过名字构建人对象的函数
func NewPerson(name string) *person {
	return &person{name: name}
}

// 将人这个对象的名字和年龄输出的函数
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

// 输出信息的函数
func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

// 定义一个WalkRun的接口
type WalkRun interface {
	Walk() //定义内部的走的函数
	Run()  //定义内部的跑的函数
}
