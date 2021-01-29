package interceptor

import (
	
	"github.com/gin-gonic/gin"
)

//GeneralInterceptor - call this medthod to add Interceptor
func GeneralInterceptor(c *gin.Context){

	token :=c.Query("token")
	if token=="1234"{
		c.Next()
	} else{
	 c.JSON(400 , gin.H{"status":"error"})
	 c.Abort()
	}
}