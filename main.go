package main

import (
	"enterbj_print/controllers"
	"github.com/gin-gonic/gin"
)

func main()  {

	router := gin.Default()

	// 首页、登录页、注册页
	router.GET("/")
	router.GET("/register")
	router.GET("/login")


	userInfo := router.Group("/user",controllers.CarLoginedMiddelWare())
	{
		// 用户个人首页
		userInfo.GET("/")

	}


	router.Run(":8000")

}

