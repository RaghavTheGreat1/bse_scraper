package main

import (
	"fmt"

	"github.com/RaghavTheGreat1/bse_scraper/routes"
	"github.com/RaghavTheGreat1/bse_scraper/services"
	"github.com/gin-gonic/gin"
)

func main() {
	services.PlaywrightContext = services.InitializePlaywright()

	r := gin.Default()
	routes.InitializeRoutes(r)

	port := ":8080"
	err := r.Run(port)

	if err != nil {
		fmt.Println("Error starting server at port", port)
		panic(err)
	}

}
