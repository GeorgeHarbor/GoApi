package main

import (
	"github.com/GeorgeHarbor/GoApi/controllers"
	"github.com/gin-gonic/gin"
)

func main()  {
 	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsGetAll)
  r.GET("/posts/:id", controllers.PostsGet)
	r.Run() 
}
