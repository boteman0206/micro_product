package goRpcAdvance

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"micro_product/golangRpc/golangRpcProto"
	"micro_product/micro_proto/hello"
	"net"
)

/**
使用证书验证的rpc调用
*/
func RegisterHello11() {

	//serverCrt := `D:\micre_project\micro_product\golangRpc\goRpcAdvance\server.crt`
	//serverKey := `D:\micre_project\micro_product\golangRpc\goRpcAdvance\server.key`

	serverCrt := `./golangRpc/goRpcAdvance/server.crt`
	serverKey := `./golangRpc/goRpcAdvance/server.key`

	// 使用证书
	creds, err := credentials.NewServerTLSFromFile(serverCrt, serverKey)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	hello.RegisterHelloServiceServer(grpcServer, new(golangRpcProto.HelloServiceImpl))

	lis, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)

}
