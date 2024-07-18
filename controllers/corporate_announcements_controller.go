package controllers

import (
	"encoding/json"
	"io"
	"log"

	"github.com/RaghavTheGreat1/bse_scraper/services"
	"github.com/gin-gonic/gin"
)

func GetCorporateAnnouncementsController(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatalf("could not read request body: %v", err)
	}

	requestBodyJson := make(map[string]interface{})

	err = json.Unmarshal(body, &requestBodyJson)

	if err != nil {
		log.Fatalf("could not unmarshal request body: %v", err)
	}

	pages := int(requestBodyJson["pages"].(float64))

	corporateAnnouncements, extractedPages, err := services.ExtractCorporateAnnouncements(services.PlaywrightContext, pages)

	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"status":  "success",
		"pages":   extractedPages,
		"count":   len(corporateAnnouncements),
		"message": corporateAnnouncements,
	})
}
