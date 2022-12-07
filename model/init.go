package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// 链接数据库操作
func Database(connstring string) {
	db, err := gorm.Open("mysql", connstring) //gorm.Open()需要学习
	if err != nil {
		fmt.Println(err)
		panic("数据库链接失败")
	}
	fmt.Println("数据库连接成功")
	db.LogMode(true)             //输出日志              //需要认识
	if gin.Mode() == "release" { //需要认识
		db.LogMode(false) //不输出日志
	}
	db.SingularTable(true)                       //表明不加s		//需要认识
	db.DB().SetMaxIdleConns(20)                  //设置连接池		//需要认识
	db.DB().SetMaxOpenConns(100)                 //最大连接数		//需要认识
	db.DB().SetConnMaxLifetime(time.Second * 30) //连接时间			//需要认识
	DB = db
	migration() //数据库迁移，没表则新建表
}
