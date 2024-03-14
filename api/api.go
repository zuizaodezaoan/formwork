package api

import (
	"context"

	"github.com/zuizaodezaoan/formwork/consul"
	"github.com/zuizaodezaoan/formwork/model"
	"github.com/zuizaodezaoan/formwork/nacos"
)

func Init(serverName string, str ...string) error {
	var err error
	//err = nacos.InitConfig()
	err = nacos.InitNacos()

	_, err = consul.InitRegisterServer(context.Background(), "user_srv.g5")
	if err != nil {
		return err
	}

	nacos.Enroll()

	if err != nil {
		return err
	}
	for _, val := range str {
		switch val {
		case "mysql":
			err = model.InitMysql(serverName)
		}

	}
	return err
}
