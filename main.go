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
	r := gin.Default()
	r.POST("/posts", controllers.CreatePosts)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run()
}
