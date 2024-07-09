package main

import (
	"bufio"
	"errors"
	"fmt"
	"gotour/ch26/server"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//创建一个自定义的用户对象
	client, err := DialHTTP("tcp", "localhost:1234")
	if err != nil { //如果有问题,直接输出错误信息
		log.Fatal("dialing:", err)
	}

	args := server.Args{A: 7, B: 8}                    //创建请求参数对象
	var reply int                                      //创建返回参数
	err = client.Call("MathService.Add", args, &reply) //调用服务器相加函数
	if err != nil {                                    //如果调用出错输出错误信息
		log.Fatal("MathService.Add error:", err)
	}
	fmt.Printf("MathService.Add: %d+%d=%d", args.A, args.B, reply) //正常情况输出信息
}

// DialHTTP connects to an HTTP RPC server at the specified network address
// listening on the default HTTP RPC path.
func DialHTTP(network, address string) (*rpc.Client, error) {
	return DialHTTPPath(network, address, rpc.DefaultRPCPath) //直接调用后面自定义的函数
}

// DialHTTPPath connects to an HTTP RPC server
// at the specified network address and path.
func DialHTTPPath(network, address, path string) (*rpc.Client, error) {
	var err error                           //创建错误信息对象
	conn, err := net.Dial(network, address) //执行dial函数
	if err != nil {                         //如果出错直接返回错误信息
		return nil, err
	}
	io.WriteString(conn, "GET "+path+" HTTP/1.0\n\n") //将获得的信息写入conn中

	// Require successful HTTP response
	// before switching to RPC protocol.
	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: "GET"}) //读取服务器返回的内容
	connected := "200 Connected to JSON RPC"                                            //创建一个连接字符串
	if err == nil && resp.Status == connected {                                         //如果当前没有出错
		return jsonrpc.NewClient(conn), nil //直接返回正确的信息
	}
	if err == nil { //如果错误为空
		err = errors.New("unexpected HTTP response: " + resp.Status) //返回错误的代码
	}
	conn.Close()              //关闭连接
	return nil, &net.OpError{ //返回错误结构体
		Op:   "dial-http",
		Net:  network + " " + address,
		Addr: nil,
		Err:  err,
	}
}
