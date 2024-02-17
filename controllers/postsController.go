package controllers

import (
	"github.com/GeorgeHarbor/GoApi/initializers"
	"github.com/GeorgeHarbor/GoApi/models"
	"github.com/gin-gonic/gin"
)


func PostsCreate (c *gin.Context) {
  // Get data off req body
  var body struct{
    Title string
    Body string
  }

  c.Bind(&body)
  
  // Create a post
  post := models.Post{Title: body.Title, Body: body.Body}

  result := initializers.DB.Create(&post)

  if result.Error != nil{
    c.Status(400)
    return 
  }

  
  // Return it
  c.JSON(200, gin.H{
    "post": post,
    })
	}

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

func PostsGet (c *gin.Context){
  // Get id from URL
  id := c.Param("id")

  // Get the post

  var post models.Post
  result := initializers.DB.First(&post, id)

  // Respond
  if result.Error != nil {
    c.Status(404)
    return
  }
  
  c.JSON(200, gin.H{
    "post": post,
  })
}

func PostsUpdate (c *gin.Context){
  // Get values from request body
  id := c.Param("id")
  
  var body struct {
    Title string
    Body string
  }
  
  c.Bind(&body)
  
  // Get the post
  var post models.Post
  initializers.DB.First(&post, id)

  // Update the post 
  result := initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body,})

  // Return the new post
  if result.Error != nil {
    c.Status(400)
    return
  }

  c.JSON(200, gin.H{
    "post": post,
  })
}

