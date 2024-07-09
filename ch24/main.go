package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(Stringify([]MyInt{1, 2, 3}))

	p := Point{1, 2, 3}
	ScaleAndPrint(p)
}

// 定义一个int类型的新类型
type MyInt int

// 定义一个转换字符串的函数
func (i MyInt) String() string {
	return fmt.Sprintf("%d:%d", i, i)
}

// 创建一个用于字符串后面加字符串的函数
func Stringify[T fmt.Stringer](s []T) (ret []string) {
	for _, v := range s { //逐一的添加元素
		ret = append(ret, v.String())
	}
	return ret
}

// Scale returns a copy of s with each element multiplied by c.
func Scale[S ~[]E, E constraints.Integer](s S) S {
	r := make(S, len(s))  //创建缓冲区
	for i, v := range s { //逐一的进行计算
		r[i] = v * 2
	}
	return r
}

type Point []int32 //创建一个缓存区
// 用于输出字符串的函数
func (p Point) String() string {
	var b strings.Builder //创建缓冲区对象
	b.Grow(len(p))        //根据p的长度设置b的长度
	for _, v := range p { //逐一写入其中
		b.WriteString(fmt.Sprint(v))
		b.WriteString(",")
	}
	return b.String()
}

// ScaleAndPrint doubles a Point and prints it.
func ScaleAndPrint(p Point) {
	r := Scale(p)           //测出其大小
	fmt.Println(r.String()) //输出
}
