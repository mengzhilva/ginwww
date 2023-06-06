package router


import (
	"github.com/gin-gonic/gin"
	"ginwww/configs"
	"strconv"
	"ginwww/controllers"
)

func InitRouter() {
	
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/", (&controllers.Index{}).Index) //有模板
	r.GET("/api", (&controllers.Index{}).Api) //无模板json返回
	r.GET("/adduser", (&controllers.Index{}).AddUser) //链接数据库并插入数据
	r.GET("/getuser", (&controllers.Index{}).GetUser) //
	
	r.NoRoute(func(c *gin.Context) {
    	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.Run(":"+strconv.Itoa(configs.HttpPort)) // 监听并在 配置参数 上启动服务
}