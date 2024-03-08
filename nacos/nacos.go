package nacos

import (
	"errors"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var ConfigClient config_client.IConfigClient

//func InitConfig() error {
//	// 初始化Viper
//	viper.SetConfigName("nacos")                                   // 配置文件名（不带扩展名）
//	viper.SetConfigType("yaml")                                    // 配置文件类型
//	viper.AddConfigPath("/Users/chenhaoqi/go/src/formwork/nacos/") // 配置文件路径
//
//	err := viper.ReadInConfig()
//	if err != nil {
//		log.Println("读取配置文件失败", err.Error())
//		return errors.New("读取配置文件失败" + err.Error())
//	}
//
//	err = viper.UnmarshalKey("nacos", &config.Usersrv.Nacos)
//	fmt.Println("配置信息：*********************", config.Usersrv)
//	if err != nil {
//		log.Println("反序列化失败", err.Error())
//		return errors.New("反序列化失败" + err.Error())
//
//	}
//	return nil
//}

func InitNacos() error {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "785c73e1-7df9-4ed4-ae03-d72b121ebe46", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   uint64(8848),
		},
	}

	var err error
	ConfigClient, err = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return errors.New("连接nacos失败" + err.Error())
	}

	return nil
}

func NacosConfig(serverName string) (string, error) {
	content, err := ConfigClient.GetConfig(vo.ConfigParam{
		DataId: "user_srv.g5",
		Group:  "json"})
	if err != nil {
		return "", err
	}
	return content, nil
}
