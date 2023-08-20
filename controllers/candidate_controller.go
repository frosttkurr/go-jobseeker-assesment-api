package controllers

import (
	"go-jobseeker-assesment-api/helpers"
	"go-jobseeker-assesment-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexCandidates(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	page_size, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	candidates, page_total_data, total_data, err := models.GetCandidates(page, page_size)
	if err == nil {
		switch {
		case total_data <= 0:
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
				"result":          candidates,
				"page_total_data": page_total_data,
				"total_data":      total_data,
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

	status_code, err := models.InsertCandidate(request_candidate)
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

	status_code, err := models.UpdateCandidate(helpers.StringToNumber(id), request_candidate)
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

	status_code, err := models.DeleteCandidate(helpers.StringToNumber(id))
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
