package api

import (
	"main/db"
	"main/interceptor"
	"main/model"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func setupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func login(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		var queryUser model.User
		if err := db.GetDB().First(&queryUser, "username= ?", user.Username).Error; err != nil {
			c.JSON(400, gin.H{"status": "nok", "err": err})
		} else if ceckPasswordHash(user.Password, queryUser.Password) == false {
			c.JSON(400, gin.H{"status": "nok", "error": "invalid user"})
		} else {
			token := interceptor.JwtSign(queryUser)
			c.JSON(200, gin.H{"result": "ok", "token": token})

		}
	} else {
		c.JSON(400, gin.H{"result": "untable to blind data"})
	}
}

func register(c *gin.Context) {

	var user model.User
	if c.ShouldBind(&user) == nil {

		user.Password, _ = hashPassword(user.Password)
		user.CratedAt = time.Now()

		if err := db.GetDB().Create(&user).Error; err != nil {
			c.JSON(200, gin.H{"result": "nook", "error": err})
		} else {
			c.JSON(200, gin.H{"result": "ok", "data": user})
		}
	} else {
		c.JSON(200, gin.H{"status": "unable "})
	}

}

func ceckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}
