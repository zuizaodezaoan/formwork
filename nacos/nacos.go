package nacos

import (
	"errors"
	"fmt"
	"log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

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
		DataId: serverName,
		Group:  config.Usersrv.Nacos.Group})
	if err != nil {
		return "", err
	}
	return content, nil
}
