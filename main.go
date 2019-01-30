package main

import (
	"enterbj_print/controllers"
	"enterbj_print/models"
	"github.com/gin-gonic/gin"
)

func init() {

	//	链接mongodb数据库
	models.LoadMongoDBInfo("192.168.1.125:27017", "", "")
}

func main() {

	router := gin.Default()

	// 首页、登录页、注册页
	router.GET("/")
	router.GET("/register")
	router.GET("/login")

	userInfo := router.Group("/user", controllers.CarLoginedMiddelWare())
	{
		// 用户个人首页
		userInfo.GET("/")

	}

	router.Run(":8000")

}
