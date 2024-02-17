package controllers

import (
	"github.com/GeorgeHarbor/GoApi/initializers"
	"github.com/GeorgeHarbor/GoApi/models"
	"github.com/gin-gonic/gin"
)


func PostsGetAll (c *gin.Context){
  // Get the posts
  var posts []models.Post
  result := initializers.DB.Find(&posts)

  // Respond
  if result.Error != nil {
    c.Status(404)
    return
  }
  
  c.JSON(200, gin.H{
    "posts": posts,
  })
}

