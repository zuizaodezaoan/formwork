package model

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/zuizaodezaoan/formwork/config"
	"github.com/zuizaodezaoan/formwork/nacos"
)

var DB *gorm.DB
var err error

func InitMysql(serverName string) error {
	nacosConfig, err := nacos.NacosConfig(serverName)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(nacosConfig), &config.Usersrv)

	fmt.Println("数据库连接配置", config.Usersrv.Mysql)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Usersrv.Mysql.Number, config.Usersrv.Mysql.Password, config.Usersrv.Mysql.Host, config.Usersrv.Mysql.Port, config.Usersrv.Mysql.DbName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger_web
			ParameterizedQueries:      true,        // Don't include params in the SQL logs
			Colorful:                  true,        // Disable color
		},
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}
	fmt.Println("321", DB)
	return nil
}

func Transaction(txc func(tx *gorm.DB) error) {
	tx := DB.Begin()
	err = txc(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}
