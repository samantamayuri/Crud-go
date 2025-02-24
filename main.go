package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samantamayuri/Crud-go/controllers"
	"github.com/samantamayuri/Crud-go/initializers"
	"github.com/samantamayuri/Crud-go/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()

}

func main() {

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.POST("/posts", middleware.Authorize, controllers.CreatePosts)
	r.GET("/posts", middleware.Authorize, controllers.GetPosts)
	r.GET("/posts/:postId", controllers.GetPost)
	r.PUT("/posts/:postId", middleware.Authorize, controllers.UpdatePost)
	r.DELETE("/posts/:postId", middleware.Authorize, controllers.DeletePost)

	r.POST("/posts/:postId/comments", controllers.CreateComment)
	r.PUT("/posts/:postId/comments/:commentId", controllers.UpdateComment)
	r.DELETE("/posts/:postId/comments/:commentId", controllers.DeleteComment)

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.Authorize, controllers.ValidateUser)
	r.Run()
}
