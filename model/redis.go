package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	config2 "github.com/zuizaodezaoan/formwork/config"
	"github.com/zuizaodezaoan/formwork/nacos"
)

//var Redis *redis.Client

func InitRedis(serverName string, hand func(cli *redis.Client) error) error {
	nacosConfig, err := nacos.NacosConfig(serverName)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(nacosConfig), &config2.Usersrv)
	if err != nil {
		return err
	}

	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config2.Usersrv.Redis.Host, config2.Usersrv.Redis.Port),
		DB:   0,
	})
	defer cli.Close()

	err = hand(cli)
	if err != nil {
		return err
	}

	return nil
}

func GetByKey(ctx context.Context, serviceName, key string) (string, error) {
	var data string
	var err error
	err = InitRedis(serviceName, func(cli *redis.Client) error {
		data, err = cli.Get(ctx, key).Result()
		return err
	})
	if err != nil {
		return "", err

	}
	return data, nil
}

func GetByRight(ctx context.Context, serviceName, key string) (bool, error) {
	var data int64
	var err error
	err = InitRedis(serviceName, func(cli *redis.Client) error {
		data, err = cli.Exists(ctx, key).Result()
		return err
	})
	if err != nil {
		return false, err
	}
	if data > 0 {
		return true, err
	}

	return false, nil
}

func GetMessage(ctx context.Context, serverName string, key string, val interface{}, duration time.Duration) error {
	return InitRedis(serverName, func(cli *redis.Client) error {
		return cli.Set(ctx, key, val, duration).Err()
	})
}
