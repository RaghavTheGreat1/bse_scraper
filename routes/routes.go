package routes

import (
	"github.com/RaghavTheGreat1/bse_scraper/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {

	r.GET("/corporate_announcements", controllers.GetCorporateAnnouncementsController)
}
