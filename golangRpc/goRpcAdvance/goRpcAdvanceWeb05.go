package goRpcAdvance

import (
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"micro_product/golangRpc/golangRpcProto"
	"micro_product/micro_proto/hello"
	"net/http"
	"strings"
)

const PORT = "9003"

/**
提供rpc和http服务共存
*/
func RegisterHello15() {

	grpcServer := grpc.NewServer()

	// http接口提供http服务
	mux := GetHTTPServeMux()

	// grpc服务架设在http之上
	hello.RegisterHelloServiceServer(grpcServer, &golangRpcProto.HelloServiceImpl{})

	// 和服务没有关系
	reflection.Register(grpcServer)

	// todo 不能使用这个，应为必须要使用http2.0才可以
	//err := http.ListenAndServe(":"+PORT,
	//	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
	//			grpcServer.ServeHTTP(w, r)
	//		} else {
	//			mux.ServeHTTP(w, r)
	//		}
	//		return
	//	}),
	//)

	// 使用这个才可以
	err := http.ListenAndServe(":9011",
		h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		}), &http2.Server{}),
	)

	if err != nil {
		log.Fatal(err.Error())
	}

}

func GetHTTPServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello-test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello: go-grpc-example"))
	})

	return mux
}
