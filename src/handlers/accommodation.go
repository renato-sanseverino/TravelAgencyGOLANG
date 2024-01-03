package handlers

import (
	"net/http"
	"travelagency/src/domain"
	"travelagency/src/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)


func PostAccommodation(c *gin.Context) {	
	var payload domain.Accommodation

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var pool pgxpool.Pool; // TODO: retornar o pool de conexões
	rep := repository.NewAccommodationRepository(&pool);
	err := rep.Insert(c, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Accommodation created successfully", "Accommodation": "ac"})
}
