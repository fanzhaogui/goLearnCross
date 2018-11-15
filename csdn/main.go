package main

import (
	"csdn/controllers"
	"github.com/gin-gonic/gin"
)

// gin-gonic/gin

func main()  {

	// 全局设置环境
	gin.SetMode(gin.DebugMode)
	// or
	gin.SetMode(gin.ReleaseMode)

	// 获取路由
	/*r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "pong",
		})
	})

	r.Run()*/

	// 构建restfull api
	router := gin.Default()

	v1 := router.Group("api/v1/userinfo")
	{
		//v1.POST("/", CreateUser)
		//v1.GET("/", FetchAllUsers)
		v1.GET("/:id", controllers.GetUser)
		//v1.PUT("/:id", UpdateUser)
		//v1.DELETE("/:id", DeleteUser)
	}

	router.Run()
}
