package nacos

import (
	"encoding/json"
	"log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"

	"github.com/zuizaodezaoan/formwork/config"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("nacos.yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Println("读取配置文件失败")
	}

	err = v.UnmarshalKey("nacos", &config.Nacoss)
	if err != nil {
		log.Println("反序列化失败")
	}

	NacosConfig()
}

func NacosConfig() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.Nacoss.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		NotLoadCacheAtStart: true,
		LogDir:              config.Nacoss.LogDir,
		CacheDir:            config.Nacoss.CacheDir,
		LogLevel:            config.Nacoss.LogLevel,
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: config.Nacoss.Host,
			Port:   uint64(config.Nacoss.Port),
		},
	}

	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: config.Nacoss.DataId,
		Group:  config.Nacoss.Group})

	err := json.Unmarshal([]byte(content), &config.Nacoss)
	if err != nil {
		return
	}
}
