package controllers

import (
	"go-jobseeker-assesment-api/models"

	"github.com/gin-gonic/gin"
)

func IndexCandidates(c *gin.Context) {
	candidates, err := models.GetCandidates()

	if err == nil {
		switch {
		case len(candidates) <= 0:
			c.JSON(204, gin.H{
				"meta": gin.H{
					"status":  204,
					"message": "No data found",
				},
			})
		default:
			c.JSON(200, gin.H{
				"meta": gin.H{
					"status":  200,
					"message": "Successfully retrieve data",
				},
				"result": candidates,
			})
		}
	} else {
		c.JSON(500, gin.H{
			"meta": gin.H{
				"status":  500,
				"message": "Failed to retrieve data",
			},
		})
	}
}

func ShowCandidates(c *gin.Context) {
	id := c.Param("id")
	candidate, err := models.FindCandidates(id)

	if err == nil {
		switch candidate.Candidate_ID {
		case 0:
			c.JSON(204, gin.H{
				"meta": gin.H{
					"status":  204,
					"message": "No data found",
				},
			})
		default:
			c.JSON(200, gin.H{
				"meta": gin.H{
					"status":  200,
					"message": "Successfully retrieve data",
				},
				"result": candidate,
			})
		}
	} else {
		c.JSON(500, gin.H{
			"meta": gin.H{
				"status":  500,
				"message": "Failed to retrieve data",
			},
		})
	}
}
