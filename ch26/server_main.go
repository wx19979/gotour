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
	//注册一个服务
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
