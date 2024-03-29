package consul

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"

	config2 "github.com/zuizaodezaoan/formwork/config"
	"github.com/zuizaodezaoan/formwork/nacos"
)

var (
	ConsulClient *api.Client
	SrvId        string
	err          error
)

func InitRegisterServer(ctx context.Context, serverName string) (string, error) {
	//使用默认配置
	config := api.DefaultConfig()
	nacosConfig, err := nacos.NacosConfig(serverName)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal([]byte(nacosConfig), &config2.Usersrv)
	if err != nil {
		return "", err
	}
	log.Println("获取nacos配置信息consul客户端=========", config2.Usersrv)
	//配置consul的连接地址
	config.Address = fmt.Sprintf("%s:%d", config2.Usersrv.Consul.Host, config2.Usersrv.Consul.Port)

	//示例化客户端
	ConsulClient, err = api.NewClient(config)

	if err != nil {
		zap.S().Panic(err.Error())
	}
	log.Println("************")
	config2.Usersrv.Host = getHostIp()

	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", config2.Usersrv.Host, config2.Usersrv.Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//健康检查,检查我们注册的微服务
	Registration := api.AgentServiceRegistration{}
	Registration.Address = config2.Usersrv.Host
	Registration.Port = config2.Usersrv.Port
	Registration.Name = config2.Usersrv.Mysql.DbName
	Registration.Tags = config2.Usersrv.Tags
	Registration.ID = fmt.Sprintf("%s", uuid.NewV4())
	SrvId = Registration.ID
	Registration.Check = check

	err = ConsulClient.Agent().ServiceRegister(&Registration)
	if err != nil {
		zap.S().Panic(err.Error())
	}

	return "", err
}

func getHostIp() string {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get current host ip err: ", err)
		return ""
	}
	var ip string
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				break
			}
		}
	}
	return ip
}

func GetServer(serverName string) (string, int, error) {
	name, i, _ := ConsulClient.Agent().AgentHealthServiceByName(serverName)
	if name != "passing" {
		log.Printf("获取nacos服务发现失败!")
	}

	var Address string
	var Port int

	for _, v := range i {
		Address = v.Service.Address
		Port = v.Service.Port
	}
	log.Println("Address: ", Address, Port)

	return Address, Port, nil
}
