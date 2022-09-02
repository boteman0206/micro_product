package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"micro_product/config"
	"micro_product/controller"
	"micro_product/micro_proto/pc"
	"micro_product/services"
	"net"
	"net/http"
)

func main() {

	// 读取配置文件
	config.InitConfig()

	go func() {
		profPort := fmt.Sprintf(":%d", config.ConfigRes.Ser.PprofPort)
		err := http.ListenAndServe(profPort, nil)
		if err != nil {
			return
		}
	}() //pprof性能分析

	flag.Parse()

	// 初始化db和redis服务
	services.SetupDB()
	defer services.CloseDB()

	// 启动rpc服务
	grpcServer := grpc.NewServer()

	pc.RegisterDcProductServer(grpcServer, new(controller.DcProduct))

	port := fmt.Sprintf(":%d", config.ConfigRes.Ser.HttpPort)
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)

}
