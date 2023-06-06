package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rawHttp "net/http"
	"ginwww/modules"
	"ginwww/utls"
)

type Index struct {
	Base
}

//展示到模板
func (controller *Index) Index(c *gin.Context) {
	fmt.Println(rawHttp.StatusOK)
	c.HTML(rawHttp.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
}

//json返回数据
func (controller *Index) Api(c *gin.Context) {
	
	c.JSON(rawHttp.StatusOK, controller.success("ok", nil))
}

//添加数据到数据库
func (controller *Index) AddUser(c *gin.Context) {
	db := utls.GetDB() 
	user := modules.User{}
	user.Id = 1
	user.Name = "aa"
	db.Create(&user)
	c.JSON(rawHttp.StatusOK, controller.success("ok", nil))
}
//查询数据
func (controller *Index) GetUser(c *gin.Context) {
	db := utls.GetDB() 
	user := modules.User{}
	db.Raw("SELECT id, name FROM go_user WHERE id = ?", 1).Scan(&user)

	c.JSON(rawHttp.StatusOK, controller.success("ok", user))
}