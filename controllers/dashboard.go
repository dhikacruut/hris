package controllers

import (
	"net/http"

	"github.com/dhikacruut/hris/db"
	"github.com/dhikacruut/hris/models"
	"github.com/gin-gonic/gin"
)

func GetEmpCount(c *gin.Context) {
	var count int64
	var employee models.Employee
	err := db.DB.Model(&employee).Where("resigndate IS NULL").Count(&count).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to count employee",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
