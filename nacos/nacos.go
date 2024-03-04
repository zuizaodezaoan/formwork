package nacos

import (
	"errors"
	"fmt"
	"log"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/util"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"

	"github.com/zuizaodezaoan/formwork/config"
)

var ConfigClient config_client.IConfigClient

func InitConfig() error {
	// 初始化Viper
	viper.SetConfigName("nacos")                                   // 配置文件名（不带扩展名）
	viper.SetConfigType("yaml")                                    // 配置文件类型
	viper.AddConfigPath("/Users/chenhaoqi/go/src/formwork/nacos/") // 配置文件路径

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("读取配置文件失败", err.Error())
		return errors.New("读取配置文件失败" + err.Error())
	}

	err = viper.UnmarshalKey("nacos", &config.Usersrv.Nacos)
	fmt.Println("配置信息：*********************", config.Usersrv)
	if err != nil {
		log.Println("反序列化失败", err.Error())
		return errors.New("反序列化失败" + err.Error())

	}
	return nil
}

func InitNacos() error {
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.Usersrv.Nacos.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		NotLoadCacheAtStart: true,
		LogDir:              config.Usersrv.Nacos.LogDir,
		CacheDir:            config.Usersrv.Nacos.CacheDir,
		LogLevel:            config.Usersrv.Nacos.LogLevel,
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: config.Usersrv.Nacos.Host,
			Port:   uint64(config.Usersrv.Nacos.Port),
		},
	}

	// 创建 Nacos 服务发现客户端
	namingClient, _ := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr:      config.Usersrv.Nacos.Host,
				Port:        uint64(config.Usersrv.Nacos.Port),
				ContextPath: "/config",
			},
		},
		"clientConfig": clientConfig,
	})

	var err error
	ConfigClient, err = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return errors.New("连接nacos失败" + err.Error())
	}

	nacos := vo.RegisterInstanceParam{
		Ip:          "10.2.171.28",                        // 服务实例的 IP 地址
		Port:        8081,                                 // 服务运行的端口号
		ServiceName: "demo.go",                            // 要注册的服务名称
		Weight:      10,                                   // 服务实例的权重（用于负载均衡）
		Enable:      true,                                 // 服务实例的启用状态
		Healthy:     true,                                 // 服务实例的健康状态
		Ephemeral:   true,                                 // 表示实例是否是临时的（不可达时将被删除）
		Metadata:    map[string]string{"idc": "shanghai"}, // 与服务实例关联的元数据
		ClusterName: "cluster-a",                          // 服务实例所属的集群名称（默认为 "DEFAULT"）
		GroupName:   "group-a",                            // 服务实例所属的组名称
	}
	_, err = namingClient.RegisterInstance(nacos)
	if err != nil {
		return errors.New("注册服务失败" + err.Error())
	}

	// SelectOneHealthyInstance return one instance by WRR strategy for load balance
	// And the instance should be health=true,enable=true and weight>0
	_, err = namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
	})

	// Subscribe key = serviceName+groupName+cluster
	// Note: We call add multiple SubscribeCallback with the same key.
	err = namingClient.Subscribe(&vo.SubscribeParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
		SubscribeCallback: func(services []model.Instance, err error) {
			log.Printf("\n\n callback return services:%s \n\n", util.ToJsonString(services))
		},
	})

	return nil

}

func NacosConfig(serverName string) (string, error) {

	content, err := ConfigClient.GetConfig(vo.ConfigParam{
		DataId: serverName,
		Group:  config.Usersrv.Nacos.Group})
	if err != nil {
		return "", err
	}
	return content, nil
}
