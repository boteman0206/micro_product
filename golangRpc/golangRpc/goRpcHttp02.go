package golangRpc

import (
	"log"
	"net/http"
	"net/rpc"
)

/**
参考文档： https://darjun.github.io/2020/05/08/godailylib/rpc/
*/
type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

/**
RegisterHello05版本没有效果，换成这种方式 架设在http协议上
*/
func RegisterHello06() {
	arith := new(Arith)
	rpc.Register(arith)

	// 这里必须要添加这个  使用http://127.0.0.1:1234/debug/rpc可以查看到注册的http方法
	rpc.HandleHTTP()
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("serve error:", err)
	}
}
