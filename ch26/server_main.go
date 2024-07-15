// RPC也就是远程过程调用,是分布式系统中不同节点的调用的方式,属于C/S模式。
// RPC由客户端发起,调用服务端的方法进行通信,然后服务端把结果返回给客户端
// RPC核心由通信协议和序列化
// 序列化和反序列化是一种把传输内容编码和解码的方式常见的编码方式有JSON、Protobuf等
// RPC调用流程:
// 1,客户端调用客户端存根同时把参数传给客户端存根
// 2,客户端存根将参数打包编码通过系统调用发送到服务端
// 3,客户端本地系统发送信息到服务器
// 4,服务器系统将信息发送到服务端存根
// 5,服务端存根解析信息,也就是解码
// 6,服务端存根调用真正的服务端程序
// 7,服务端处理后通过同样方式,把结果再返回给客户端
// func (t *T)MethodName(argType T1,replyType *T2)error
// 第一个参数argType是调用者(客户端)提供的
// 第二个参数replyType是返回给调用者结果,必须是指针类型
// Call方法参数的作用
// 1,调用远程方法的名字,点前面是注册的服务的名称,点后面的部分是该服务的方法
// 2,客户端为了调用远程方法提供的参数,示例中是args
// 3,为接收远程方法返回的结果,必须是一个指针
package main

import (
	"gotour/ch26/server"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//注册一个服务对象(1,服务的名称2,具体的服务对象)
	rpc.RegisterName("MathService", new(server.MathService))

	//注册一个path，用于提供基于http的json rpc服务
	http.HandleFunc(rpc.DefaultRPCPath, func(rw http.ResponseWriter, r *http.Request) {
		conn, _, err := rw.(http.Hijacker).Hijack()
		if err != nil {
			log.Print("rpc hijacking ", r.RemoteAddr, ": ", err.Error())
			return
		}
		//创建一个rpc的连接成功的消息字符串
		var connected = "200 Connected to JSON RPC"
		io.WriteString(conn, "HTTP/1.0 "+connected+"\n\n") //在conn写入这些信息
		jsonrpc.ServeConn(conn)                            //根据这个链接提供服务
	})

	l, e := net.Listen("tcp", ":1234") //设置监听对象
	if e != nil {                      //如果有问题直接返回错误信息
		log.Fatal("listen error:", e)
	}

	http.Serve(l, nil) //换成http的服务
}
