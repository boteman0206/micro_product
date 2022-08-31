package golangRpcProto

import (
	"google.golang.org/grpc"
	"log"
	"micro_product/micro_proto/pc"
	"net"
)

/**

grpc的流服务的入参：HelloService_ChannelServer

*/
func (p *HelloServiceImpl) ChannelOneWay(stream pc.HelloService_ChannelOneWayServer) error {
	//服务端在循环中接收客户端发来的数据，如果遇到 io.EOF 表示客户端流被关闭，如果函数退出表示服务端流关闭
	//生成返回的数据通过流发送给客户端，双向流数据的发送和接收都是完全独立的行为。需要注意的是，发送和接收的操作并不需要一一对应，用户可以根据真实场景进行组织代码。

	for i := 0; i < 5; i++ {
		recv, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Println("客户端信息：", recv)
	}

	// 服务端最后一条消息发送
	err := stream.SendAndClose(&pc.StringDto{Value: "关闭"})
	if err != nil {
		return err
	}
	return nil

}

func RegisterHello09() {
	grpcServer := grpc.NewServer()
	pc.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)

}
