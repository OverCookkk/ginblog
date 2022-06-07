package model

import (
	"fmt"
	"ginblog/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// db的入口文件

var db *gorm.DB
var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName)
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Printf("mysql conn err ,%s, dsn:%s\n", err, dsn)
	}
	// db.SingularTable(true)
	// 数据库迁移
	db.AutoMigrate(&User{}, &Category{}, &Article{})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)                  // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                 // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // SetConnMaxLifetime 设置了连接可复用的最大时间。

	// db.Close()
}
