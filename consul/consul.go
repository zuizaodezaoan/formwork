package consul

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"

	config2 "github.com/zuizaodezaoan/formwork/config"
)

var (
	ConsulClient *api.Client
	SrvId        string
	err          error
)

func InitRegisterServer(ctx context.Context, serverName string) (string, error) {
	//使用默认配置
	config := api.DefaultConfig()

	//配置consul的连接地址
	config.Address = fmt.Sprintf("%s:%d", config2.Usersrv.Consul.Host, config2.Usersrv.Consul.Port)

	//示例化客户端
	ConsulClient, err = api.NewClient(config)

	if err != nil {
		zap.S().Panic(err.Error())
	}

	//right, err := model.GetByRight(ctx, serverName, "consul: node: index")
	//if err != nil {
	//	log.Println("233333333", err)
	//	return "", err
	//}

	//key, err := model.GetByKey(ctx, serverName, "consul: node: index")
	//if err != nil {
	//	log.Println("44444444444000000", err)
	//
	//	return "", err
	//}
	//index, _ := strconv.Atoi(key)
	//

	//if right {
	//key, err := model.GetByKey(ctx, serverName, "consul: node: index")
	//if err != nil {
	//	log.Println("44444444444000000", err)
	//
	//	return "", err
	//}
	//
	//index, _ := strconv.Atoi(key)
	//log.Println("oooooooooooooooo")
	//err = model.RedisIndexAdd(ctx, serverName, "consul: node: index")
	//if err != nil {
	//	return "", err
	//}
	//
	//err = model.GetMessage(ctx, serverName, "consul: node: index", index+1, 0)
	//if err != nil {
	//	log.Println("5555555555", err)
	//
	//	return "", err
	//}

	//}
	//err = model.GetMessage(ctx, serverName, "consul: node: index", 1, 0)

	//log.Println("hhhhhhhhhhhhhhhh")
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
	name, i, _ := ConsulClient.Agent().AgentHealthServiceByName(config2.Usersrv.Mysql.DbName)
	if name != "passing" {
		log.Printf("获取nacos服务发现失败!")
	}

	var Address string
	var Port int

	for _, v := range i {
		Address = v.Service.Address
		Port = v.Service.Port
	}

	return Address, Port, nil
}
