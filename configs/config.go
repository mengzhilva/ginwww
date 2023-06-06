package configs


import "fmt"
import "github.com/go-ini/ini"

var HttpPort int


// MaxIdle 最大空闲连接数
var MaxIdle int

// MaxActive 最大活跃连接数
var MaxActive int

// IdleTimeout 空闲连接超时时间
var IdleTimeout int

// MysqlHost mysql主机
var MysqlHost string

// MysqlPort mysql端口
var MysqlPort int

// MysqlUser mysql用户名
var MysqlUser string

// MysqlPassword mysql密码
var MysqlPassword string

// MysqlDbName mysql数据库名
var MysqlDbName string

// MysqlCharset mysql字符集
var MysqlCharset string

// MysqlMaxOpenConn mysql最大连接数
var MysqlMaxOpenConn int

// MysqlMaxIdleConn mysql最大空闲连接数
var MysqlMaxIdleConn int


func init(){
	cfg, err := ini.Load("configs/config.ini")
	if err != nil {
		fmt.Println("读取ini文件失败")
	}
	serverConfig := cfg.Section("server")

	HttpPort = serverConfig.Key("http_port").MustInt()
	// 读取mysql配置文件
	mysqlConfig := cfg.Section("mysql")
	MysqlHost = mysqlConfig.Key("host").String()
	MysqlPort = mysqlConfig.Key("port").MustInt()
	MysqlUser = mysqlConfig.Key("username").String()
	//mysql密码中包含#
	MysqlPassword = mysqlConfig.Key("password").String()
	//
	MysqlDbName = mysqlConfig.Key("dbName").String()
	MysqlCharset = mysqlConfig.Key("charset").String()
	MysqlMaxOpenConn = mysqlConfig.Key("maxOpenConn").MustInt()
	MysqlMaxIdleConn = mysqlConfig.Key("maxIdleConn").MustInt()
}

