package goRpcAdvance

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"micro_product/golangRpc/golangRpcProto"
	"micro_product/micro_proto/hello"
	"net/http"
	"strings"
)

/**
提供rpc和http服务共存
在上面15的例子上，加上了gin或者echo的路由处理
*/
func RegisterHello16() {

	grpcServer := grpc.NewServer()

	//// grpc服务架设在http之上
	hello.RegisterHelloServiceServer(grpcServer, &golangRpcProto.HelloServiceImpl{})
	//
	//// 和服务没有关系
	reflection.Register(grpcServer)

	// http接口提供http服务
	engine := gin.Default()

	engine.Any("/hello", func(c *gin.Context) {

		var data = make(map[string]interface{}, 0)

		err := c.Bind(&data)
		if err != nil {
			return
		}
		c.JSON(200, data)
	})

	engine.Use(func(r *gin.Context) {
		// 判断协议是否为http/2  // 判断是否是grpc
		if r.Request.ProtoMajor == 2 && strings.Contains(r.GetHeader("Content-Type"), "application/grpc") {
			// 按grpc方式来请求
			grpcServer.ServeHTTP(r.Writer, r.Request)
			// 不要再往下请求了,防止继续链式调用拦截器
			r.Abort()
			return
		} else {
			// 当作普通api
			r.Next()
		}

	})

	handler := h2c.NewHandler(engine, &http2.Server{})

	h := &http.Server{
		Addr:    ":9002",
		Handler: handler,
	}

	err := h.ListenAndServe()
	if err != nil {
		return
	}
}
