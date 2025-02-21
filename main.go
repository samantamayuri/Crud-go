package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samantamayuri/Crud-go/controllers"
	"github.com/samantamayuri/Crud-go/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()

}

func main() {

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.POST("/posts", controllers.CreatePosts)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:postId", controllers.GetPost)
	r.PUT("/posts/:postId", controllers.UpdatePost)
	r.DELETE("/posts/:postId", controllers.DeletePost)

	r.POST("/posts/:postId/comments", controllers.CreateComment)
	r.PUT("/posts/:postId/comments/:commentId", controllers.UpdateComment)
	r.DELETE("/posts/:postId/comments/:commentId", controllers.DeleteComment)
	r.Run()
}
