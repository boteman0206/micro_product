

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


-- 自签名证书的使用 + grpc
linux生成证书的命令
$ openssl genrsa -out server.key 2048
$ openssl req -new -x509 -days 3650 -subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io"  -key server.key -out server.crt

$ openssl genrsa -out client.key 2048
$ openssl req -new -x509 -days 3650     -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io"  -key client.key -out client.crt
以上命令将生成 server.key、server.crt、client.key 和 client.crt 四个文件。
其中以. key 为后缀名的是私钥文件，需要妥善保管。以. crt 为后缀名是证书文件，也可以简单理解为公钥文件，并不需要秘密保存。
在 subj 参数中的 /CN=server.grpc.io 表示服务器的名字为 server.grpc.io，在验证服务器的证书时需要用到该信息。
--  使用例子在：goRpcAdvance01文件中



-- 根证书的使用 + grpc
为了避免证书的传递过程中被篡改，可以通过一个安全可靠的根证书分别对服务器和客户端的证书进行签名。这样客户端或服务器在收到对方的证书后可以通过根证书进行验证证书的有效性。
根证书的生成方式和自签名证书的生成方式类似：
$ openssl genrsa -out ca.key 2048
$ openssl req -new -x509 -days 3650    -subj "/C=GB/L=China/O=gobook/CN=github.com"  -key ca.key -out ca.crt
然后是重新对服务器端证书进行签名：
$ openssl req -new -subj "/C=GB/L=China/O=server/CN=server.io" -key server.key  -out server.csr
$ openssl x509 -req -sha256  -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 -in server.csr   -out server.crt
签名的过程中引入了一个新的以. csr 为后缀名的文件，它表示证书签名请求文件。在证书签名完成之后可以删除. csr 文件。
然后在客户端就可以基于 CA 证书对服务器进行证书验证：
-- 使用例子在 goRpcAdvance02文件中 这个还没有跑通，待后续验证




-- Token 认证
前面讲述的基于证书的认证是针对每个 gRPC 连接的认证。gRPC 还为每个 gRPC 方法调用提供了认证支持，这样就基于用户 Token 对不同的方法访问进行权限管理。

