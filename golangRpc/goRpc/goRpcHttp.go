package goRpc

import (
	"fmt"
	"io"
	"micro_product/micro_proto"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**


todo 这里为啥通过http调用不了
在http上构建rpc服务

Go 语言内在的 RPC 框架已经支持在 Http 协议上提供 RPC 服务。但是框架的 http 服务同样采用了内置的 gob 协议，并且没有提供采用其它协议的接口，因此从其它语言依然无法访问的。
在前面的例子中，我们已经实现了在 TCP 协议之上运行 jsonrpc 服务，并且通过 nc 命令行工具成功实现了 RPC 方法调用。现在我们尝试在 http 协议上提供 jsonrpc 服务。

新的 RPC 服务其实是一个类似 REST 规范的接口，接收请求并采用相应处理流程：
*/

type HelloServiceJsonHttp struct{}

func (p *HelloServiceJsonHttp) Hello(request string, reply *string) error {
	*reply = request
	fmt.Println("RegisterHello05: ", reply)
	return nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	var data = []byte{}
	read, err := r.Body.Read(data)
	if err != nil {
		return
	}
	fmt.Println("入参：", read, " data：", data)
	fmt.Fprintf(w, "Hello!")
}

func RegisterHello05() {

	err := rpc.RegisterName(micro_proto.HelloServiceName, new(HelloServiceJsonHttp))
	if err != nil {
		return
	}

	/**
	RPC 的服务架设在 “/jsonrpc” 路径，在处理函数中基于 http.ResponseWriter 和 http.Request 类型的参数构造一个 io.ReadWriteCloser 类型的 conn 通道。
	然后基于 conn 构建针对服务端的 json 编码解码器。最后通过 rpc.ServeRequest 函数为每次请求处理一次 RPC 方法调用
	*/
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			fmt.Println("rpc.ServeRequest(jsonrpc.NewServerCodec(conn)) error: ", err.Error())
			return
		}
		fmt.Fprintf(conn, "json http response")
	})

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":1234", nil)
}
