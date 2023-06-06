package utls

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ginwww/configs"
)

// Db 数据库连接
var Db *gorm.DB

// Init 初始化数据库连接,设置连接池
func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", configs.MysqlUser, configs.MysqlPassword, configs.MysqlHost, configs.MysqlPort, configs.MysqlDbName, configs.MysqlCharset)
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
	} else {
		sqlDB, _ := Db.DB()
		sqlDB.SetMaxOpenConns(configs.MysqlMaxOpenConn)//设置数据库连接池最大连接数
		sqlDB.SetMaxIdleConns(configs.MysqlMaxIdleConn)//连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	}
}

//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过utls.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return Db
}
