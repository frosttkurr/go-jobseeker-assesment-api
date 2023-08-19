package controllers

import (
	"go-jobseeker-assesment-fullstack/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var response gin.H
	candidates, err := models.GetCandidates()

	if err == nil {
		response = gin.H{
			"meta": gin.H{
				"status":  200,
				"message": "Successfully get data",
			},
			"result": candidates,
		}
		c.JSON(200, response)
	} else {
		response = gin.H{
			"meta": gin.H{
				"status":  500,
				"message": "Failed to get data",
			},
		}
		c.JSON(500, response)
	}
}
