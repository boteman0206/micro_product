

//  golangRpc模块： 测试golang自带的rpc测试
使用golang自带的rpc框架的rpc服务端（调用端在micro_api项目中）
包含了普通的rpc和jsonrpc的基本使用

其中http的rpc指的是将rpc架构在http协议的基础上来进行调用



// golangRpcProto模块： grpc+protobuf配合使用的服务端
使用golang+proto的测试文件，对原始golang知道的rpc的升级 以及gRPC 流的使用

安装工具protobuf
首先是安装官方的 protoc 工具，可以从 https://github.com/google/protobuf/releases 下载。

然后是安装针对 Go 语言的代码生成插件
可以通过 go get github.com/golang/protobuf/protoc-gen-go 命令安装。


生成go代码  其中 go_out 参数告知 protoc 编译器去加载对应的 protoc-gen-go 工具，然后通过该工具生成代码，生成代码放到当前目录。最后是一系列要处理的 protobuf 文件的列表。
$ protoc --go_out=. hello.proto   注意：这样只会生成message相关的pb文件


生成grpc文件 使用 protoc-gen-go 内置的 gRPC 插件生成 gRPC 代码： gRPC 插件会为服务端和客户端生成不同的接口：
$ protoc --go_out=plugins=grpc:. hello.proto




// goRpcAdvance模块：主要展示的是证书的使用和一些protobuf的扩展
证书认证： gRPC 建立在 HTTP/2 协议之上，对 TLS 提供了很好的支持。 我们前面章节中 gRPC 的服务都没有提供证书支持，因此客户端在连接服务器中通过 grpc.WithInsecure()


linux生成证书的命令
$ openssl genrsa -out server.key 2048
$ openssl req -new -x509 -days 3650 -subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io"  -key server.key -out server.crt

$ openssl genrsa -out client.key 2048
$ openssl req -new -x509 -days 3650     -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io"  -key client.key -out client.crt
以上命令将生成 server.key、server.crt、client.key 和 client.crt 四个文件。
其中以. key 为后缀名的是私钥文件，需要妥善保管。以. crt 为后缀名是证书文件，也可以简单理解为公钥文件，并不需要秘密保存。
在 subj 参数中的 /CN=server.grpc.io 表示服务器的名字为 server.grpc.io，在验证服务器的证书时需要用到该信息。