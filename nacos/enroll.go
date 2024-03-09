package nacos

import (
	"fmt"
	"log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var Client naming_client.INamingClient
var err error

func Enroll() {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create naming client
	Client, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	instance, err := Client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "10.2.171.28",
		Port:        8081,
		ServiceName: "demo.go",
		GroupName:   "group-a",
		ClusterName: "cluster-a",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
	})
	if err != nil {
		fmt.Println("失败", instance)
		return
	}

}

func GetNacosSrv() (string, int, error) {
	service, err := Client.GetService(vo.GetServiceParam{
		Clusters:    nil,
		ServiceName: "demo.go",
		GroupName:   "group-a",
	})
	if err != nil {
		return "", 0, err
	}

	var host string
	var port int

	for _, v := range service.Hosts {
		host = v.Ip
		port = int(v.Port)

	}
	log.Println("Host:========== ", host, port)
	return host, port, nil
}
