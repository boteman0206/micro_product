package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"micro_product/config"
	"micro_product/controller"
	"micro_product/micro_proto/pc"
	"micro_product/services"
	"net"
	"net/http"
	"strings"
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

	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系，有些grpc的工具依赖这个
	reflection.Register(grpcServer)

	port := fmt.Sprintf(":%d", config.ConfigRes.Ser.HttpPort)
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)

}

/**
http和grpc共存的服务使用
*/
func main01() {

	// http接口提供http服务
	mux := GetHTTPServeMux()

	grpcServer := grpc.NewServer()

	// grpc服务架设在http之上
	pc.RegisterDcProductServer(grpcServer, new(controller.DcProduct))

	reflection.Register(grpcServer)

	err := http.ListenAndServe(":9001",
		h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		}), &http2.Server{}),
	)

	if err != nil {
		log.Fatal(err)
	}

}

func GetHTTPServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello-test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("this is  my test ....")
		w.Write([]byte("eddycjy: go-grpc-example"))
	})

	return mux
}
