package main

import (
	"net/http"
	"travelagency/src/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin gonic API",
		})
	})

	rGroup := router.Group("/api")
	rGroup.GET("/itineraries", handlers.GetItineraries)
	rGroup.POST("/itineraries", handlers.PostItinerary)
	rGroup.PATCH("/itineraries/:id", handlers.PatchItinerary)
	rGroup.DELETE("/itineraries/:id", handlers.DeleteItinerary)
	rGroup.GET("/clients", handlers.GetClients)
	rGroup.POST("/clients", handlers.PostClient)
	rGroup.PATCH("/clients/:id", handlers.PatchClient)
	rGroup.DELETE("/clients/:id", handlers.DeleteClient)
	rGroup.POST("/hotels", handlers.PostAccommodation)

	router.Use(cors.Default())
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
