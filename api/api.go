package api

import (
	"github.com/redis/go-redis/v9"

	"github.com/zuizaodezaoan/formwork/model"
	"github.com/zuizaodezaoan/formwork/nacos"
)

func Init(serverName string, str ...string) error {
	var err error
	err = nacos.InitConfig()
	err = nacos.InitNacos()

	nacos.Enroll()

	if err != nil {
		return err
	}
	for _, val := range str {
		switch val {
		case "mysql":
			err = model.InitMysql(serverName)
		case "redis":
			err = model.InitRedis(serverName, func(cli *redis.Client) error {
				return err
			})
		}
	}
	return err
}
