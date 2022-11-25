package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"mall_srvs/order_srv/model"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

// create database mall_order_srv CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci
func main() {
	dsn := "root:12312321@tcp(42.377:3306)/mall_order_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.ShoppingCart{}, &model.OrderInfo{}, &model.OrderGoods{})
	if err != nil {
		panic(err)
	}
}
