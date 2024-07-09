package main

//缓存已经计算的结果
var cache = map[int]int{}

// 斐波那契函数
func Fibonacci(n int) int {
	//获取缓存的信息
	if v, ok := cache[n]; ok {
		return v
	}
	//初始化结果寄存器
	result := 0
	//根据情况进行计算斐波那契
	switch {
	case n < 0: //如果n小于0直接置为零
		result = 0
	case n == 0: //如果n等于0也置为零
		result = 0
	case n == 1: //n等于1结果直接就是1
		result = 1
	default: //其他情况调用斐波那契函数
		result = Fibonacci(n-1) + Fibonacci(n-2)
	}
	cache[n] = result //临时寄存之前的计算结果
	return result     //最后返回结果
}
