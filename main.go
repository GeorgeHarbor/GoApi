package main

import (
	"github.com/GeorgeHarbor/GoApi/controllers"
	"github.com/gin-gonic/gin"
)

func main()  {
 	r := gin.Default()
	r.GET("/posts", controllers.PostsGetAll)
	r.Run() 
}
