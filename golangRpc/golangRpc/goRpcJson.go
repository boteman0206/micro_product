package golangRpc

import (
	"fmt"
	"log"
	"micro_product/micro_proto"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
标准库的 RPC 默认采用 Go 语言特有的 gob 编码，因此从其它语言调用 Go 语言实现的 RPC 服务将比较困难。
在互联网的微服务时代，每个 RPC 以及服务的使用者都可能采用不同的编程语言，因此跨语言是互联网时代 RPC 的一个首要条件。
得益于 RPC 的框架设计，Go 语言的 RPC 其实也是很容易实现跨语言支持的。


Go 语言的 RPC 框架有两个比较有特色的设计：一个是 RPC 数据打包时可以通过插件实现自定义的编码和解码；另一个是 RPC 建立在抽象的 io.ReadWriteCloser 接口之上的，
我们可以将 RPC 架设在不同的通讯协议之上。这里我们将尝试通过官方自带的 net/rpc/jsonrpc 扩展实现一个跨语言的 RPC。
*/

type HelloServiceJson struct{}

func (p *HelloServiceJson) Hello(request string, reply *string) error {
	*reply = request
	fmt.Println("RegisterHello04: ", request)
	return nil
}

func RegisterHello04() {

	rpc.RegisterName(micro_proto.HelloServiceName, new(HelloServiceJson))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		/**
		使用jsonrpc代码
		*/
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
