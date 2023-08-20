package controllers

import (
	"go-jobseeker-assesment-api/helpers"
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
	candidate, err := models.FindCandidate(helpers.StringToNumber(id))

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

	if err_input := c.ShouldBindJSON(&request_candidate); err_input != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status":  http.StatusBadRequest,
				"message": err_input.Error(),
			},
		})
		return
	}

	err, status_code := models.InsertCandidate(request_candidate)
	if err != nil {
		c.JSON(status_code, gin.H{
			"meta": gin.H{
				"status":  status_code,
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(status_code, gin.H{
		"meta": gin.H{
			"status":  status_code,
			"message": "Successfully insert new data",
		},
	})
}

func UpdateCandidate(c *gin.Context) {
	var request_candidate models.T_Candidate
	id := c.Param("id")

	if err_input := c.ShouldBindJSON(&request_candidate); err_input != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status":  http.StatusBadRequest,
				"message": err_input.Error(),
			},
		})
		return
	}

	err, status_code := models.UpdateCandidate(helpers.StringToNumber(id), request_candidate)
	if err != nil {
		c.JSON(status_code, gin.H{
			"meta": gin.H{
				"status":  status_code,
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(status_code, gin.H{
		"meta": gin.H{
			"status":  status_code,
			"message": "Successfully update data",
		},
	})
}

func DeleteCandidate(c *gin.Context) {
	id := c.Param("id")

	err, status_code := models.DeleteCandidate(helpers.StringToNumber(id))
	if err != nil {
		c.JSON(status_code, gin.H{
			"meta": gin.H{
				"status":  status_code,
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(status_code, gin.H{
		"meta": gin.H{
			"status":  status_code,
			"message": "Successfully delete data",
		},
	})
}
