package main

import "fmt"

func main() {
	var s string           //创建一个字符串
	fmt.Printf("%p\n", &s) //输出字符串的地址
	s = "张三"               //创建一个字符串
	fmt.Println(s)         //输出字符串具体内容

	var sp *string   //创建一个字符串的指针
	sp = new(string) //申请一个字符串内存空间
	*sp = "飞雪无情"     //设置数据内容
	fmt.Println(*sp)

	pp := NewPerson("飞雪无情", 20)                    //根据构造函数创建一个人的对象
	fmt.Println("name为", pp.name, ",age为", pp.age) //输出结果

	m := map[string]int{"张三": 18} //创建一个map对象
	fmt.Println(m)                //输出信息

}

//创建人对象的函数
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
