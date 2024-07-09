// 某一个包中的函数想要是公有函数需要将函数名首字母大写,才可以在不同的包中调用
// 函数名称首字母是小写的属于私有函数只能在同一个包中调用
package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := sum(1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println()
	fmt.Println("飞雪")
	fmt.Println("飞雪", "无情")

	fmt.Println(sum1(1, 2))
	fmt.Println(sum1(1, 2, 3))
	fmt.Println(sum1(1, 2, 3, 4))
	//可以直接在函数里面声明一个函数名然后在后面进行调用
	sum2 := func(a, b int) int {
		return a + b
	}
	//直接调用之前定义过的sum2的函数
	fmt.Println(sum2(1, 2))

	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	age := Age(25)
	age.String()
	age.Modify()
	age.String()

	sm := Age.String //方法赋值给变量,也就是方法表达式
	sm(age)          //通过变量,要穿一个接收者进行调用所以要传一个age
}

// 多值返回函数(并且含有函数返回值命名)
func sum(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	sum = a + b
	err = nil
	return
}

// 可以传入任意多个参数,获取参数可以直接遍历参数列表即可(可变参数函数,类型前面加三个点,其实就是该类型的切片)
func sum1(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}

// 函数作为编程的一等公民也可以作为返回值进行返回
func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//函数与方法的区别:方法必须要有一个接收者,这个接收者是一个类型,这样方法就和这个类型绑定在一起,称为这个类型的方法
//函数属于一个包,而方法是属于一个类型

type Age uint

// 该方法是属于Age这一类型的方法
func (age Age) String() {
	fmt.Println("the age is", age)
}

// 方法也可以是指针类型
func (age *Age) Modify() {
	*age = Age(30)
}
