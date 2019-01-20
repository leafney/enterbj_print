package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {

	router := gin.Default()

	//v1 := router.Group("api/v1")
	//{
	//	// routes.Pong(v1)
	//	//routes.Index(v1)
	//
	//}


	router.Run(":8000")

}

