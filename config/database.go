package config

import (
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 数据库
type DataBase struct {
	Connection string
	Host       string
	Port       string
	DataBase   string
	UserName   string
	Password   string
	Prefix     string
}

func GetDatabaseConfig() (db DataBase) {
	db.Connection = "mysql"
	db.Host     = "127.0.0.1"
	db.Port     = "3306"
	db.DataBase = "go"
	db.UserName = "root"
	db.Password = "memedama"
	return db
}

func GetDB() *gorm.DB {
	var DB *gorm.DB
	var dsn string

	var err error
	conf := GetDatabaseConfig()

	dsn = conf.UserName + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/"
	dsn += conf.DataBase + "?charset=utf8&parseTime=True&loc=Asia%2FShanghai"

	DB, err = gorm.Open(conf.Connection, dsn)
	if err != nil {
		log.Fatalf("failed to connect database %v ", err)
	}

	return DB
}
