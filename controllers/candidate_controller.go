package controllers

import (
	"go-jobseeker-assesment-api/models"
	"net/http"

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

func ShowCandidate(c *gin.Context) {
	id := c.Param("id")
	candidate, err := models.FindCandidate(id)

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

func CreateCandidate(c *gin.Context) {
	var request_candidate models.T_Candidate

	if err := c.ShouldBindJSON(&request_candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		})
		return
	}

	if err := models.InsertCandidate(request_candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"status":  http.StatusOK,
			"message": "Successfully insert new data",
		},
	})
}
