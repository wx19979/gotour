// 创建一个服务对象包
package server

//创建一个计算服务结构体
type MathService struct {
}

// 定义一个参数的结构体
type Args struct {
	A, B int //参数里面含有A,B两个操作数
}

//创建一个给该服务结构体相加的函数
func (m *MathService) Add(args Args, reply *int) error {
	*reply = args.A + args.B //将A,B的相加的结果放到返回的对象中
	return nil               //返回结果是没有错
}
