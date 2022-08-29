package goRpc

import (
	"log"
	"micro_product/micro_proto"
	"net"
	"net/rpc"
)

/**
对上一个版本的优化
*/
//const HelloServiceName = "path/to/pkg.HelloService" // 放到proto里面

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(micro_proto.HelloServiceName, svc)
}

//注册hello的rpc方法
func RegisterHello02() {

	//HelloService 类型的对象注册为一个 RPC 服务：

	RegisterHelloService(new(HelloService))

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
