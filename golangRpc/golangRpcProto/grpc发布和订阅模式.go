package golangRpcProto

/**
在发布和订阅模式中，由调用者主动发起的发布行为类似一个普通函数调用，
而被动的订阅者则类似 gRPC 客户端单向流中的接收者。现在我们可以尝试基于 gRPC 的流特性构造一个发布和订阅系统。


这里先不做介绍  https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-04-grpc.html#443-grpc-%E6%B5%81
*/

import (
	"context"
	"fmt"
	"github.com/moby/moby/pkg/pubsub"
	"google.golang.org/grpc"
	"log"
	"micro_product/micro_proto/pc"
	"net"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

//然后是实现发布方法和订阅方法：
func (p *PubsubService) Publish(ctx context.Context, arg *pc.StringDto) (*pc.StringDto, error) {
	fmt.Println("Publish: ", arg.GetValue())
	p.pub.Publish(arg.GetValue())

	return &pc.StringDto{}, nil
}

//然后是实现发布方法和订阅方法：
func (p *PubsubService) Subscribe(arg *pc.StringDto, stream pc.PubsubService_SubscribeServer) error {
	fmt.Println("Subscribe: ", arg.GetValue())
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			// 接收数据是string，并且key是以arg为前缀
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	// 遍历服务器的chan，并将其中的信息发送给订阅客户端
	for v := range ch {
		if err := stream.Send(&pc.StringDto{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

/**
发布订阅的服务端
*/
func RegisterHello10() {

	service := NewPubsubService()
	grpcServer := grpc.NewServer()
	pc.RegisterPubsubServiceServer(grpcServer, service)

	lis, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)

}

/**
这个是docker项目的一个例子： 发布订阅是一个常见的设计模式，开源社区中已经存在很多该模式的实现。其中 docker 项目中提供了一个 pubsub 的极简实现，下面是基于 pubsub 包实现的本地发布订阅代码：
*/
func DockerPublishSubscribeExample() {

	//其中 pubsub.NewPublisher 构造一个发布对象
	p := pubsub.NewPublisher(100*time.Millisecond, 10)

	//p.SubscribeTopic() 可以通过函数筛选感兴趣的主题进行订阅。
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})
	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})

	go p.Publish("hi")
	go p.Publish("golang: https://golang.org")
	go p.Publish("docker: https://www.docker.com/")
	time.Sleep(1)

	go func() {
		fmt.Println("golang topic:", <-golang)
	}()
	go func() {
		fmt.Println("docker topic:", <-docker)
	}()

	// 阻塞
	for true {

	}
}
