package main

import (
	"github.com/GeorgeHarbor/GoApi/initializers"
	"github.com/GeorgeHarbor/GoApi/models"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDB()
}

func main() {
  initializers.DB.AutoMigrate(&models.Post{})
}
