package goRpcAdvance

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"
	"micro_product/micro_proto/hello"
	"net"
)

/**
使用token认证的grpc的服务端

*/

type Authentication struct {
	User     string
	Password string
}

type grpcServerT struct {
	auth *Authentication
}

func (p *grpcServerT) Hello(ctx context.Context, in *hello.StringDto) (*hello.StringDto, error) {
	if err := p.auth.Auth(ctx); err != nil {
		return nil, err
	}
	return &hello.StringDto{Value: " token 验证返回: " + in.Value}, nil
}

/**
详细地认证工作主要在 Authentication.Auth 方法中完成。
首先通过 metadata.FromIncomingContext 从 ctx 上下文中获取元信息，然后取出相应的认证信息进行认证。如果认证失败，则返回一个 codes.Unauthenticated 类型地错误。
*/
func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var appid string
	var appkey string

	if val, ok := md["user"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	if appid != a.User || appkey != a.Password {
		return grpc.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}

/**
token验证的服务端
*/
func RegisterHello12() {

	server := grpc.NewServer()

	t := &grpcServerT{
		auth: &Authentication{
			User:     "gopher",
			Password: "password",
		},
	}
	hello.RegisterAuthServiceServer(server, t)

	lis, err := net.Listen("tcp", ":8001")

	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)

}
