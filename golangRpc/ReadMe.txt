

// golangRpc模块： 测试golang自带的rpc测试
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

