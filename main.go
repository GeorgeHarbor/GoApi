package main

import (
	"github.com/GeorgeHarbor/GoApi/controllers"
	"github.com/GeorgeHarbor/GoApi/initializers"
	"github.com/gin-gonic/gin"
)

func init()  {
 initializers.LoadEnvVariables()
  initializers.ConnectToDB()
}

func main()  {
 	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
  r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsGetAll)
  r.GET("/posts/:id", controllers.PostsGet)
  r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run() 
}
