package common

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	ServerPort string

	RedisHost string
	RedisPort string
	RedisPassword string

	DbSource string
	DbHost string
	DbName string
	DbUser string
	DbPassword string
	DbPort string
	DbMaxIdleConns int
	DbMaxOpenConns int
)

func init() {
	file, error := ini.Load("common/MarvelousBlog-Backend-config.ini")
	if error != nil {
		fmt.Println("配置文件读入异常")
	}

	loadServer(file)
	loadRedis(file)
	loadDatabase(file)
}

func loadDatabase(file *ini.File) {
	DbSource = file.Section("Database").Key("DbSource").String()
	DbHost = file.Section("Database").Key("DbHost").String()
	DbName = file.Section("Database").Key("DbName").String()
	DbUser = file.Section("Database").Key("DbUser").String()
	DbPort = file.Section("Database").Key("DbPort").String()
	DbPassword = file.Section("Database").Key("DbPassword").String()
	DbMaxIdleConns, _ = file.Section("Database").Key("DbMaxIdleConns").Int()
	DbMaxOpenConns, _ = file.Section("Database").Key("DbMaxOpenConns").Int()
}

func loadRedis(file *ini.File) {
	RedisHost = file.Section("Redis").Key("RedisHost").String()
	RedisPort = file.Section("Redis").Key("RedisPort").String()
	RedisPassword = file.Section("Redis").Key("RedisPassword").String()
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	ServerPort = file.Section("server").Key("ServerPort").String()
}