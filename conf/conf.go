package conf

import (
	"TodoList/model"
	"fmt"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string //debug模式
	HttpPort   string //服务器3000端口
	Db         string //数据库
	DbHost     string //数据库主机地址
	DbPort     string //数据库端口号
	DbUser     string //数据库账户
	DbPassword string //数据库密码
	DbName     string //数据库名
)

func Init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		fmt.Println("读取配置文件错误，请检查文件路径")
	}
	LoadServer(file)
	LoadMysql(file)
	//拼接数据库链接  strings.Join 需要学习 user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
	path := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	// fmt.Println(path)
	model.Database(path)
}

// 读取服务配置信息  ini.File 需要学习
func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

// 读取数据库配置信息 ini.File 需要学习
func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
