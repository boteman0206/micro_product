package goRpcAdvance

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	serverCrt := `./golangRpc/goRpcAdvance/server.crt`
	serverKey := `./golangRpc/goRpcAdvance/server.key`

	// 首先 gRPC 是建立在 HTTP/2 版本之上，如果 HTTP 不是 HTTP/2 协议则必然无法提供 gRPC 支持。
	// 同时，每个 gRPC 调用请求的 Content-Type 类型会被标注为 "application/grpc" 类型。
	creds, err := credentials.NewServerTLSFromFile(serverCrt, serverKey)
	if err != nil {
		log.Fatal(err)
	}

	// http接口提供http服务
	mux := GetHTTPServeMux()

	// grpc服务架设在http之上
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	hello.RegisterHelloServiceServer(grpcServer, &golangRpcProto.HelloServiceImpl{})

	http.ListenAndServeTLS(":"+PORT, serverCrt, serverKey,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
			return
		}),
	)

}

func GetHTTPServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello-test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("eddycjy: go-grpc-example"))
	})

	return mux
}
