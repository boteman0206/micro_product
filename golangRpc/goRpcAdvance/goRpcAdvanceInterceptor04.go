package goRpcAdvance

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"log"
	"micro_product/golangRpc/golangRpcProto"
	"micro_product/micro_proto/hello"
	"net"
)

/**
不过 gRPC 框架中只能为每个服务设置一个截取器，因此所有的截取工作只能在一个函数中完成。
开源的 grpc-ecosystem 项目中的 go-grpc-middleware 包已经基于 gRPC 对截取器实现了链式截取器的支持。
以下是 go-grpc-middleware 包中链式截取器的简单用法
*/

//多个filter的使用
func filter2(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	/**
	函数的 ctx 和 req 参数就是每个普通的 RPC 方法的前两个参数。第三个 info 参数表示当前是对应的那个 gRPC 方法，
	第四个 handler 参数对应当前的 gRPC 方法函数。上面的函数中首先是日志输出 info 参数，然后调用 handler 对应的 gRPC 方法函数。
	要使用 filter 截取器函数，只需要在启动 gRPC 服务时作为参数输入即可：
	*/
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
			fmt.Println("filter2拦截器捕获：", err.Error())
		}
	}()

	log.Println("filter2:req ", req, " info: ", info)
	if data, ok := req.(*hello.StringDto); ok {
		fmt.Println("filter2 获取拦截器的参数： ", data.Value)
	}

	return handler(ctx, req)
}

/*
 grpc的拦截器的实现
然后服务器在收到每个 gRPC 方法调用之前，会首先输出一行日志，然后再调用对方的方法。
如果截取器函数返回了错误，那么该次 gRPC 方法调用将被视作失败处理。
因此，我们可以在截取器中对输入的参数做一些简单的验证工作。同样，也可以对 handler 返回的结果做一些验证工作。截取器也非常适合前面对 Token 认证工作。
*/
func RegisterHello14() {

	// 添加多个拦截器使用
	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		filter, filter2,
	)))

	// 流截取
	//grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	//	filter, filter2,
	//)),

	hello.RegisterHelloServiceServer(server, new(golangRpcProto.HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)

}
