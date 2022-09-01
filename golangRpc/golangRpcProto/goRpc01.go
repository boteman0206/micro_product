package golangRpcProto

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"micro_product/micro_proto/hello"
	"net"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *hello.StringDto) (*hello.StringDto, error) {

	reply := &hello.StringDto{Value: "hello:" + args.GetValue()}
	return reply, nil

}

/**

首先是通过 grpc.NewServer() 构造一个 gRPC 服务对象，
然后通过 gRPC 插件生成的 RegisterHelloServiceServer 函数注册我们实现的 HelloServiceImpl 服务。然后通过 grpcServer.Serve(lis) 在一个监听端口上提供 gRPC 服务。

*/
func RegisterHello07() {
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)

}
