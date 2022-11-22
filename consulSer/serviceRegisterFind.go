package consulSer

import (
	"context"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	myConfig "micro_product/config"
	"micro_product/micro_common/utils"
	"micro_product/micro_proto/pc"
	"net/http"
	"strings"
)

/**
注册
*/

func consulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consulCheck")
}

// 注册服务
func RegisterServer() {

	config := consulapi.DefaultConfig()

	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	fmt.Println("HealthPort: ", myConfig.ConfigRes.Ser.HealthPort)
	checkPort := myConfig.ConfigRes.Ser.HealthPort

	registration := new(consulapi.AgentServiceRegistration)
	registration.Port = myConfig.ConfigRes.Ser.HttpPort
	registration.ID = "micro_product:" + utils.GetIP() + ":" + cast.ToString(registration.Port)
	registration.Name = "micro_product"
	registration.Tags = []string{"micro_product", "v1.0"}
	registration.Address = utils.GetIP()

	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s", //check失败后30秒删除本服务
	}

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

	http.HandleFunc("/check", consulCheck)
	err = http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}

}

//获取服务
func GetConsulServices(serName string) *consulapi.AgentService {

	client, err := consulapi.NewClient(consulapi.DefaultConfig()) //非默认情况下需要设置实际的参数
	if err != nil {
		log.Println("NewClient consul err： ", err)
	}
	services, err := client.Agent().Services()
	if err != nil {
		log.Println("Services consul err： ", err)
	}

	for k, v := range services {
		if strings.HasPrefix(k, serName) {
			return v
		}
	}

	return nil
}

//获取服务
func GetInstancesById(serviceId string) {

	client, err := consulapi.NewClient(consulapi.DefaultConfig()) //非默认情况下需要设置实际的参数

	if err != nil {
		log.Println("NewClient consul err： ", err)
	}
	catalogService, _, _ := client.Catalog().Service(serviceId, "", nil)

	// 通过tag获取服务
	tags, q, err := client.Health().ServiceMultipleTags(serviceId, []string{}, true, nil)
	if err != nil {
		return
	}
	for i := range catalogService {
		service := catalogService[i]
		fmt.Println(service.ServiceID, service.ServiceName, service.Address, service.ServicePort, " ===========  ", utils.JsonToString(service))

	}
	for i := range tags {
		service := tags[i]
		fmt.Println(service.Service.Service, service.Service.Address, service.Service.Port, " ===========  ", utils.JsonToString(service), q)

	}
}

// 测试consul网关调用
//https://www.liwenzhou.com/posts/Go/name-resolving-and-load-balancing-in-grpc/#autoid-0-1-0
func ConsulTest() {

	dialContext, err := grpc.DialContext(
		context.Background(),
		"consul://127.0.0.1:8500/micro_product?tags=micro_product,v1.2", // 使用dns需要导入包  _ "github.com/mbobakov/grpc-consul-resolver"
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), // 设置负载均衡的策略
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*100)),
	)
	if err != nil {
		return
	}

	client := pc.NewDcProductClient(dialContext)

	for i := 0; i < 3; i++ {
		product, err := client.TestProduct(context.Background(), &pc.GetProductDto{
			Id:   0,
			Name: "",
			Sort: "",
		})
		if err != nil {
			return
		}

		fmt.Println(utils.JsonToString(product))

	}

}
