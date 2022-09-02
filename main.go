package main

import (
	"google.golang.org/grpc"
	"log"
	"micro_product/controller"
	"micro_product/micro_proto/pc"
	"net"
)

func main() {

	grpcServer := grpc.NewServer()

	pc.RegisterDcProductServer(grpcServer, new(controller.DcProduct))

	lis, err := net.Listen("tcp", ":8802")

	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)

}
