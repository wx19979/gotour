// 测试模块
// 1,含有单元测试代码的go文件必须以_test.go结尾
// 2,单元测试文件名_test.go前面的部分最好是被测试的函数所在的go文件的文件名
// 3,单元测试的函数名必须以Test开头,是可导出的、公开的函数
// 4,测试函数的签名必须接收一个testing.T类型的指针,并且不能返回任何值
// 5,函数名最好是Test+要测试的函数名
// 单元测试的重点:熟悉业务代码的逻辑、场景等以便尽可能地全面测试，保障代码质量
// go test -v ./ch18 单元测试
// go test -v --coverprofile=ch20.cover ./ch20 测试单元覆盖率
// go tool cover -html ch20.cover -o ch20.html 生成覆盖率的html
// go test -bench=./ch20 基准测试
// 基准测试必须以Benchmark开头,必须是可导出的,函数的签名必须接收一个指向testing.B类型的指针,且不能返回任何值
// 最后的for循环很重要,被测试代码要放到循环里。b.N是测试框架提供的表示循环次数,因需反复调用测试代码才可评估性能
package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	//预先定义的一组斐波那契数列作为测试用例
	fsMap := map[int]int{}
	fsMap[-1] = 0
	fsMap[0] = 0
	fsMap[1] = 1
	fsMap[2] = 1
	fsMap[3] = 2
	fsMap[4] = 3
	fsMap[5] = 5
	fsMap[6] = 8
	fsMap[7] = 13
	fsMap[8] = 21
	fsMap[9] = 34
	for k, v := range fsMap {
		fib := Fibonacci(k)
		if v == fib {
			t.Logf("结果正确:n为%d,值为%d", k, fib)
		} else {
			t.Errorf("结果错误：期望%d,但是计算的值是%d", v, fib)
		}
	}
}

// 用于基准测试
func BenchmarkFibonacci(b *testing.B) {
	n := 10
	b.ReportAllocs() //开启内存统计
	b.ResetTimer()   //重置计时器
	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}

// 用于并发基准测试
func BenchmarkFibonacciRunParallel(b *testing.B) {
	n := 10
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fibonacci(n)
		}
	})
}
