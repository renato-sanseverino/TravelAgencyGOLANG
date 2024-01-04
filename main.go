package main

import (
	// "os"
	"net/http"
	// "travelagency/src/utils"
	"travelagency/src/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func addItineraryHandlers(rGroup *gin.RouterGroup) {
	rGroup.GET("/itineraries", handlers.GetItineraries)
	rGroup.POST("/itineraries", handlers.PostItinerary)
	rGroup.PATCH("/itineraries/:id", handlers.PatchItinerary)
	rGroup.DELETE("/itineraries/:id", handlers.DeleteItinerary)
}

func addClientHandlers(rGroup *gin.RouterGroup) {
	rGroup.GET("/clients", handlers.GetClients)
	rGroup.POST("/clients", handlers.PostClient)
	rGroup.PATCH("/clients/:id", handlers.PatchClient)
	rGroup.DELETE("/clients/:id", handlers.DeleteClient)
}

func main() {
	godotenv.Load(".env")
	// var databaseURL = os.Getenv("DATABASE_URL")
	// var pool = utils.GetPool(databaseURL);

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin gonic API",
		})
	})

	rGroup := router.Group("/api")
	addItineraryHandlers(rGroup)
	addClientHandlers(rGroup)
	rGroup.POST("/hotels", handlers.PostAccommodation)
	rGroup.POST("/insurances", handlers.PostInsurance)

	router.Use(cors.Default())
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
