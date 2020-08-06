package config

import (
	"MarvelousBlog-Backend/common"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	ServerHost string
	ServerPort string

	RedisHost     string
	RedisPort     string
	RedisPassword string

	DbSource       string
	DbHost         string
	DbName         string
	DbUser         string
	DbPassword     string
	DbPort         string
	DbMaxIdleConns int
	DbMaxOpenConns int

	LogFilePath string
	LogFileName string
	LogLevel    string

	SaltPrefix  string
	SaltPostfix string
	JwtSecurity string
)

//从ini文件中读取所有输入变量
func init() {
	file, err := ini.Load("config/MarvelousBlog-Backend-config.ini")
	if err != nil {
		logrus.Errorf(common.SYSTEM_ERROR_LOG, "配置文件读入异常, errMsg = ", err)
	}

	loadServer(file)
	loadRedis(file)
	loadDatabase(file)
	loadLogging(file)
	loadSecurity(file)
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
	ServerHost = file.Section("server").Key("ServerHost").String()
}

func loadLogging(file *ini.File) {
	LogFilePath = file.Section("Logging").Key("LogFilePath").String()
	LogFileName = file.Section("Logging").Key("LogFileName").String()
	LogLevel = file.Section("Logging").Key("LogLevel").String()
}

func loadSecurity(file *ini.File) {
	SaltPrefix = file.Section("Security").Key("SaltPrefix").String()
	SaltPostfix = file.Section("Security").Key("SaltPostfix").String()
	JwtSecurity = file.Section("Security").Key("JwtSecurity").String()
}
