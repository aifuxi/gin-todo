package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败", err)
	}

	// 迁移 schema，相当于让GORM帮你创建表
	db.AutoMigrate(&Todo{})
}
