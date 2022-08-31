package golangRpcProto

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"micro_product/micro_proto/pc"
	"net"
	"time"
)

/**

grpc的流服务的入参：HelloService_ChannelServer
可以发现服务端和客户端的流辅助接口均定义了 Send 和 Recv 方法用于流数据的双向通信。

*/
func (p *HelloServiceImpl) Channel(stream pc.HelloService_ChannelServer) error {
	//服务端在循环中接收客户端发来的数据，如果遇到 io.EOF 表示客户端流被关闭，如果函数退出表示服务端流关闭
	//生成返回的数据通过流发送给客户端，双向流数据的发送和接收都是完全独立的行为。需要注意的是，发送和接收的操作并不需要一一对应，用户可以根据真实场景进行组织代码。

	for {
		args, err := stream.Recv()
		fmt.Println("stream: Recv ", args, err)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		format := time.Now().Format("2006-01-02 15:04:05")
		reply := &pc.StringDto{Value: "hello:" + args.GetValue() + " time: " + format}
		fmt.Println("stream reply: ", reply)
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}

}

func RegisterHello08() {
	grpcServer := grpc.NewServer()
	pc.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)

}
