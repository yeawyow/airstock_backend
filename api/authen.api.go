package api

import (
	"main/model"

	"github.com/gin-gonic/gin"
)

func setupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func login(c *gin.Context) {
	c.JSON(200, gin.H{"status": "is ok"})
}

func register(c *gin.Context) {

	var user model.User
	if c.ShouldBind(&user) == nil {
		c.JSON(200, gin.H{"status": "register ok", "data": user})
	}

}
