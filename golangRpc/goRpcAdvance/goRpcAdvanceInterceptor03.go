package goRpcAdvance

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"micro_product/golangRpc/golangRpcProto"
	"micro_product/micro_proto/hello"
	"net"
)

/**
gRPC 中的 grpc.UnaryInterceptor 和 grpc.StreamInterceptor 分别对普通方法和流方法提供了截取器的支持。我们这里简单介绍普通方法的截取器用法。
*/

//要实现普通方法的截取器，需要为 grpc.UnaryInterceptor 的参数实现一个函数：
func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	/**
	函数的 ctx 和 req 参数就是每个普通的 RPC 方法的前两个参数。第三个 info 参数表示当前是对应的那个 gRPC 方法，
	第四个 handler 参数对应当前的 gRPC 方法函数。上面的函数中首先是日志输出 info 参数，然后调用 handler 对应的 gRPC 方法函数。
	要使用 filter 截取器函数，只需要在启动 gRPC 服务时作为参数输入即可：
	*/
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
			fmt.Println("filter拦截器捕获：", err.Error())
		}
	}()

	log.Println("filter:req ", req, " info: ", info)
	if data, ok := req.(*hello.StringDto); ok {
		fmt.Println("获取拦截器的参数： ", data.Value)
	}

	return handler(ctx, req)
}

/*
 grpc的拦截器的实现
然后服务器在收到每个 gRPC 方法调用之前，会首先输出一行日志，然后再调用对方的方法。
如果截取器函数返回了错误，那么该次 gRPC 方法调用将被视作失败处理。
因此，我们可以在截取器中对输入的参数做一些简单的验证工作。同样，也可以对 handler 返回的结果做一些验证工作。截取器也非常适合前面对 Token 认证工作。
*/
func RegisterHello13() {

	// 添加拦截器
	server := grpc.NewServer(grpc.UnaryInterceptor(filter))

	hello.RegisterHelloServiceServer(server, new(golangRpcProto.HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)

}
