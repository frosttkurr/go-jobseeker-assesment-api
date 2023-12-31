package main

import (
	"go-jobseeker-assesment-api/controllers"
	"go-jobseeker-assesment-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	api := r.Group("api")
	{
		candidates := api.Group("/candidates")
		{
			candidates.GET("/", controllers.IndexCandidates)
			candidates.GET("/:id", controllers.ShowCandidate)
			candidates.POST("/", controllers.CreateCandidate)
			candidates.PUT("/:id", controllers.UpdateCandidate)
			candidates.DELETE("/:id", controllers.DeleteCandidate)
		}
	}

	r.Run()
}
