package main

import (
	"micro_product/golangRpc/golangRpc"
	_ "net/http/pprof"
)

func main() {

	// 注册hello的rpc方法
	//golangRpc.RegisterHello()

	//golangRpc.RegisterHello02()

	// 最终
	//golangRpc.RegisterHello03()

	//使用jsonrpc传输
	//golangRpc.RegisterHello04()

	golangRpc.RegisterHello05()

	// 使用http的方式调用
	//golangRpc.RegisterHello06()

}
