package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//数组的测试
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[2])
	array1 := [5]string{1: "b", 3: "d"}
	fmt.Println(array1)

	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, array[i])
	}

	for i, v := range array {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}
	//切片的测试
	slice := array[2:5]
	slice[1] = "f"
	fmt.Println(array)

	slice1 := []string{"a", "b", "c", "d", "e"}
	//切片后面添加元素
	slice2 := append(slice1, "f")
	fmt.Println(slice1, slice2)
	//map的测试
	nameAgeMap := make(map[string]int)
	nameAgeMap["飞雪无情"] = 20
	//map返回的头一个是具体元素值,第二个是返回结果是否正确
	age, ok := nameAgeMap["飞雪无情1"]
	if ok {
		fmt.Println(age)
	}
	//根据key值删除元素的操作
	delete(nameAgeMap, "飞雪无情")

	//测试for range
	nameAgeMap["飞雪无情"] = 20
	nameAgeMap["飞雪无情1"] = 21
	nameAgeMap["飞雪无情2"] = 22

	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ",Value is", v)
	}

	fmt.Println(len(nameAgeMap))
	//将字符串转为一个byte数组
	s := "Hello飞雪无情"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(s[0], s[1], s[15])                         //逐个输出数组的信息
	fmt.Println("字符串的长度:", len(s))                         //输出该数组的长度(汉字的字符占据3个字符长度)
	fmt.Println("真正字符串中字符的个数:", utf8.RuneCountInString(s)) //输出真正字符的个数

	for i, r := range s {
		fmt.Println(i, r)
	}

	aa := [3][3]int{}
	aa[0][0] = 1
	aa[0][1] = 2
	aa[0][2] = 3
	aa[1][0] = 4
	aa[1][1] = 5
	aa[1][2] = 6
	aa[2][0] = 7
	aa[2][1] = 8
	aa[2][2] = 9
	fmt.Println(aa)
}
