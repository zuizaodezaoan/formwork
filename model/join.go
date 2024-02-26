package model

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func InitMysql() {
	dsn := "root:root@tcp(127.0.0.1:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"
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
