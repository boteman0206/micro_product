package main

import (
	"micro_product/golangRpc/goRpc"
	_ "net/http/pprof"
)

func main() {

	// 注册hello的rpc方法
	//goRpc.RegisterHello()

	//goRpc.RegisterHello02()

	// 最终
	//goRpc.RegisterHello03()

	//使用jsonrpc传输
	goRpc.RegisterHello04()

}
