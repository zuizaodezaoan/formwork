package api

import (
	"github.com/zuizaodezaoan/formwork/model"
	"github.com/zuizaodezaoan/formwork/nacos"
)

func Init(serverName string, str ...string) error {
	var err error
	//err = nacos.InitConfig()
	err = nacos.InitNacos()

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
