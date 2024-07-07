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
	fmt.Println(sum2(1, 2))

	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	age := Age(25)
	age.String()
	age.Modify()
	age.String()

	sm := Age.String
	sm(age)
}

func sum(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	sum = a + b
	err = nil
	return
}

// 可以传入任意多个参数,获取参数可以直接遍历参数列表即可
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

type Age uint

// 该函数修改的只是拷贝值
func (age Age) String() {
	fmt.Println("the age is", age)
}

// 该函数因为是传的指针所以能够修改函数值
func (age *Age) Modify() {
	*age = Age(30)
}
