package controllers

import "github.com/gin-gonic/gin"

// 用户注册
func Register(c *gin.Context) {
	//获取 邮箱、用户名、密码、验证码、
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 判断邮箱是否已经存在

	// 判断用户名是否已经存在

}
