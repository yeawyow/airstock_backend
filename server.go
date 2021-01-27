package main

import ( 
	"main/api"
    "github.com/gin-gonic/gin"
)
 
func main() {
	router := gin.Default()
	//กำหนด static folder
	router.Static("/images", "./uploaded/images")
	
    
	api.Setup(router)
	router.Run(":8081")
}
   