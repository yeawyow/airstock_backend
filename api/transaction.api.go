package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupTransActionAPI(router *gin.Engine) {
	transActionAPI := router.Group("/api/v2")
	{
		transActionAPI.GET("/transaction", getTransaction)
		transActionAPI.POST("/transaction", createTransaction)
	}

}

func getTransaction(c *gin.Context) {
	c.String(http.StatusOK, "List Tranaction")
}

func createTransaction(c *gin.Context) {
	c.String(http.StatusOK, "Create transaction")
}
