package main

import (
	"github.com/GeorgeHarbor/GoApi/controllers"
	"github.com/gin-gonic/gin"
)

func main()  {
 	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
  r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsGetAll)
  r.GET("/posts/:id", controllers.PostsGet)
  r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run() 
}
