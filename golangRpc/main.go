package main

import (
	"micro_product/golangRpc/goRpcAdvance"
	_ "net/http/pprof"
)

func main() {

	// =======================golang自带rpc ============================================
	// 注册hello的rpc方法
	//golangRpc.RegisterHello()

	//golangRpc.RegisterHello02()

	// 最终
	//golangRpc.RegisterHello03()

	//使用jsonrpc传输
	//golangRpc.RegisterHello04()

	//golangRpc.RegisterHello05()

	// 使用http的方式调用
	//golangRpc.RegisterHello06()

	//=====================================rpc+proto服务端=================================

	//普通的proto调用
	//golangRpcProto.RegisterHello07()

	//proto的流服务端
	//golangRpcProto.RegisterHello08()

	//protod的单向流服务
	//golangRpcProto.RegisterHello09()

	// 发布订阅的例子，这是一体的,并没有提供grpc服务
	//golangRpcProto.DockerPublishSubscribeExample()
	// 发布订阅的服务端测试
	//golangRpcProto.RegisterHello10()

	//=================================grpc的高级使用=========================================================

	// grpc使用tls证书
	//goRpcAdvance.RegisterHello11()

	goRpcAdvance.RegisterHello12()

}
