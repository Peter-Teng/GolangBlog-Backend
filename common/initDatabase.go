package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var Db *gorm.DB
var err error

func init() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DbUser, DbPassword, DbHost, DbPort, DbName)
	Db, err = gorm.Open(DbSource, connectionString)
	if err != nil {
		fmt.Println()
	}

	//设置表名不加s
	Db.SingularTable(true)

	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	Db.DB().SetMaxIdleConns(DbMaxIdleConns)

	// SetMaxOpenConns 设置数据库连接最大打开数。
	Db.DB().SetMaxOpenConns(DbMaxOpenConns)

	// SetConnMaxLifetime 设置可重用连接的最长时间
	Db.DB().SetConnMaxLifetime(10 * time.Second)
}