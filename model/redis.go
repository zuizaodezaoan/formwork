package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"

	config2 "github.com/zuizaodezaoan/formwork/config"
	"github.com/zuizaodezaoan/formwork/nacos"
)

//var Redis *redis.Client

func InitRedis(serverName string, hand func(cli *redis.Client) error) error {
	log.Println("00000000")
	nacosConfig, err := nacos.NacosConfig(serverName)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(nacosConfig), &config2.Usersrv)
	if err != nil {
		return err
	}
	log.Println("grpc服务================================================================", config2.Usersrv)
	log.Println("233333333=============", config2.Usersrv.Redis)

	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config2.Usersrv.Redis.Host, config2.Usersrv.Redis.Port),
		DB:   0,
	})

	err = hand(cli)
	if err != nil {
		return err
	}

	defer cli.Close()

	return nil
}

func GetByKey(ctx context.Context, serviceName, key string) (string, error) {
	var data string
	var err error
	err = InitRedis(serviceName, func(cli *redis.Client) error {
		data, err1 := cli.Get(ctx, "login"+key).Result()
		if err != nil {
			return err1
		}
		fmt.Println("mmmmmmmm888888888", data)

		return nil
	})
	fmt.Println("12345678", data)
	return data, nil
}

func RedisIndexAdd(ctx context.Context, serviceName, key string) error {
	var err error
	err = InitRedis(serviceName, func(cli *redis.Client) error {
		err := cli.Incr(ctx, key).Err()
		if err != nil {
			return errors.New("自增key失败" + err.Error())
		}
		return nil
	})
	return err
}

func GetMessage(ctx context.Context, serverName string, key string, val interface{}, duration time.Duration) error {
	return InitRedis(serverName, func(cli *redis.Client) error {
		return cli.Set(ctx, key, val, duration).Err()
	})
}

func GetByRight(ctx context.Context, serviceName, key string) (bool, error) {
	var err error
	err = InitRedis(serviceName, func(cli *redis.Client) error {
		_, err = cli.Exists(ctx, key).Result()
		return err
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
