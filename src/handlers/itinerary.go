package handlers;

import (
	"net/http"
	"travelagency/src/utils"
	"travelagency/prisma/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetItineraries(c *gin.Context) {
	// var Itineraries []db.InnerItinerary

	client := utils.GetPrisma(c)

	Itineraries, err := client.Itinerary.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Itineraries})
}

func PostItinerary(c *gin.Context) {
	var payload db.InnerItinerary

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	insertedItinerary, err := client.Itinerary.CreateOne(
		db.Itinerary.ID.Set(payload.ID),
		db.Itinerary.Destination.Set(payload.Destination),
		db.Itinerary.Departure.Set(payload.Departure),
		db.Itinerary.TransportKind.Set(payload.TransportKind),
		db.Itinerary.Arrival.SetOptional(payload.Arrival),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Itinerary created successfully", "Itinerary": insertedItinerary})
}

func PatchItinerary(c *gin.Context) {
	var payload db.InnerItinerary

	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	updatedItinerary, err := client.Itinerary.FindUnique(
		db.Itinerary.ID.Equals(id.String()),
	).Update(
		db.Itinerary.Destination.Set(payload.Destination),
		db.Itinerary.Departure.Set(payload.Departure),
		db.Itinerary.TransportKind.Set(payload.TransportKind),
		db.Itinerary.Arrival.SetOptional(payload.Arrival),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Itinerary patched", "Itinerary": updatedItinerary})
}

func DeleteItinerary(c *gin.Context) {
	// TODO: utilizar o flag_removed ao invés de apagar o registro na tabela

	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	client := utils.GetPrisma(c)
	deletedItinerary, err := client.Itinerary.FindUnique(
		db.Itinerary.ID.Equals(id.String()),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Itinerary deleted successfully", "id": deletedItinerary.ID})
}
