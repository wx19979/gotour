package main

import (
	"fmt"
	"os"
)

const name = "飞雪无情"

func main() {
	//调用创建一个文件夹的函数
	os.Mkdir("tmp", 0666)

	fmt.Println("飞雪无情")
	//创建一个map对象
	m := map[int]string{}
	s := "飞雪无情"
	m[0] = s
}

// 用于创建一个字符串的函数
func newString() string {
	s := new(string)
	*s = "飞雪无情"
	return *s
}
