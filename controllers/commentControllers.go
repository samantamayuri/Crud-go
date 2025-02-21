package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/samantamayuri/Crud-go/initializers"
	"github.com/samantamayuri/Crud-go/models"
)

func CreateComment(c *gin.Context) {

	var body struct {
		Comment string `json:"comment"`
	}

	postId := c.Param("postId")

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to bind JSON"})
		return
	}

	var post models.Post

	result := initializers.DB.Find(&post, postId)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to get post"})
		return
	}

	if post.ID == 0 {
		c.JSON(400, gin.H{"error": "Post not found"})
		return
	}

	post.Comments = append(post.Comments, models.Comment{Comment: body.Comment, PostId: post.ID})

	result = initializers.DB.Save(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(200, gin.H{"message": "Comment created successfully", "comment": post.Comments})

}

func UpdateComment(c *gin.Context) {

	var body struct {
		Comment string `json:"comment"`
	}

	postId := c.Param("postId")
	commentId := c.Param("commentId")

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to bind JSON"})
		return
	}

	var comment models.Comment

	result := initializers.DB.Where("post_id = ? AND id = ?", postId, commentId).Find(&comment)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to find comment"})
		return
	}

	comment.Comment = body.Comment

	result = initializers.DB.Save(&comment)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(200, gin.H{"message": "Comment updated successfully", "comment": comment})

}

func DeleteComment(c *gin.Context) {

	postId := c.Param("postId")
	commentId := c.Param("commentId")

	var comment models.Comment

	result := initializers.DB.Where("post_id = ? AND id = ?", postId, commentId).Delete(&comment)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(200, gin.H{"message": "Comment deleted successfully"})

}
