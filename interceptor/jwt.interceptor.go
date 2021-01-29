package interceptor

import (
	//"net/http"
	//"strings"
	//"fmt"
	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtVerify(c *gin.Context){

	c.JSON(200, gin.H{"test":"ok"})
}