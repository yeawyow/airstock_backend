package main

import (
	//"main/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//กำหนด static folder
	router.Static("/images", "./uploaded/images")
	router.GET("/go", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/profile", func(c *gin.Context){
		c.JSON(200 ,gin.H{"status":"profile"})
	})
	//api.Setup(router)
	router.Run(":8081")
}
