package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/samantamayuri/Crud-go/initializers"
	"github.com/samantamayuri/Crud-go/models"
)

func CreatePosts(c *gin.Context) {

	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to bind JSON"})
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(200, gin.H{"message": "Post created successfully", "post": post})

}

func GetPosts(c *gin.Context) {

	var posts []models.Post

	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to get posts"})
		return
	}
	c.JSON(200, gin.H{"posts": posts})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to get post"})
		return
	}
	c.JSON(200, gin.H{"post": post})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to bind JSON"})
		return
	}

	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to update post"})
		return
	}
	post.Title = body.Title
	post.Body = body.Body

	result = initializers.DB.Save(&post)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(200, gin.H{"message": "Post updated successfully", "post": post})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}
