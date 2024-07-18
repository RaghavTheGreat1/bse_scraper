package routes

import (
	"github.com/RaghavTheGreat1/bse_scraper/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {

	r.POST("/corporate_announcements", controllers.GetCorporateAnnouncementsController)
}
