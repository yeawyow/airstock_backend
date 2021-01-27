package api

import (
	"main/db"

	"github.com/gin-gonic/gin"
)

// Setup - call this method to setup route
func Setup(router *gin.Engine) {
	
	db.SetupDB()
	setupAuthenAPI(router)
	setupProductAPI(router)
	setupTransActionAPI(router)

}
