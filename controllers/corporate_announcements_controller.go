package controllers

import (
	"github.com/RaghavTheGreat1/bse_scraper/services"
	"github.com/gin-gonic/gin"
)

func GetCorporateAnnouncementsController(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	corporateAnnouncements := services.ExtractCorporateAnnouncements(services.PlaywrightContext)

	c.IndentedJSON(200, gin.H{
		"status":  "success",
		"message": corporateAnnouncements,
	})
}
