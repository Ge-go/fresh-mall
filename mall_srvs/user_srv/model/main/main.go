package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"mall_srvs/user_srv/model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

// create database mall_user_srv CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci
func main() {
	dsn := "root:w821230693@tcp(121.37.232.8:3306)/mall_user_srv?charset=utf8mb4&parseTime=True&loc=Local"

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

	//err = db.AutoMigrate(&model.User{})
	//if err != nil {
	//	panic(err)
	//}

	//密码加密
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodePwd := password.Encode("admin123", options)
	pwd := fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodePwd)
	//创建一批测试数据
	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("ws_%d", i),
			Mobile:   fmt.Sprintf("1501245667%d", i),
			Password: pwd,
			BaseModel: model.BaseModel{
				UpdatedAt: time.Now(),
			},
		}
		db.Save(&user)
	}
}
