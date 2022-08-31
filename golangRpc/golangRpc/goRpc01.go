package golangRpc

import (
	"log"
	"net"
	"net/rpc"
)

/**
这里是使用golang自带的rpc测试

 注册rpc方法
	其中 Hello 方法必须满足 Go 语言的 RPC 规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个 error 类型，同时必须是公开的方法。
	然后就可以将 HelloService 类型的对象注册为一个 RPC 服务：
*/
type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

//注册hello的rpc方法
func RegisterHello() {

	//HelloService 类型的对象注册为一个 RPC 服务：
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
