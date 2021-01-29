package api

import (
	"main/interceptor"
	"github.com/gin-gonic/gin"
)


func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		   productAPI.POST("/product",interceptor.JwtVerify, getProduct)
		   productAPI.GET("/product", createProduct)
	}
	
}

func getProduct(c *gin.Context){
	c.JSON(200 , gin.H{"status":"product ok"})
}

func createProduct(c *gin.Context){
	c.JSON(200 , gin.H{"status":"createproduct ok"})
}