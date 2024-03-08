package consul

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"

	config2 "github.com/zuizaodezaoan/formwork/config"
	"github.com/zuizaodezaoan/formwork/model"
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

	right, err := model.GetByRight(ctx, serverName, "没有索引")
	if err != nil {
		return "", err
	}

	if right {
		key, err := model.GetByKey(ctx, serverName, "没有索引")
		if err != nil {
			return "", err
		}

		index, _ := strconv.Atoi(key)

		err = model.GetMessage(ctx, serverName, "没有索引", index+1, 0)
		if err != nil {
			return "", err
		}

	}

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
