package model

import (
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/zuizaodezaoan/formwork/config"
	"github.com/zuizaodezaoan/formwork/nacos"
)

var Redis *redis.Client

func InitRedis(serverName string) error {
	nacosConfig, err := nacos.NacosConfig(serverName)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(nacosConfig), &config.Usersrv)

	Redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Usersrv.Redis.Host, config.Usersrv.Redis.Port),
		DB:   0,
	})

	return nil
}
