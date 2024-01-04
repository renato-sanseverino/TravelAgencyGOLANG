package handlers

import (
	"os"
	"net/http"
	"travelagency/src/domain"
	"travelagency/src/repository"
	"travelagency/src/utils"
	"github.com/gin-gonic/gin"
)

func PostInsurance(c *gin.Context) {
	var databaseURL = os.Getenv("DATABASE_URL")   // passar através do MIDDLEWARE   utils.appMiddleware()
	var pool = utils.GetPool(databaseURL);        // passar através do MIDDLEWARE   utils.appMiddleware()

	var payload domain.Insurance

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rep := repository.NewInsuranceRepository(pool)
	err := rep.Insert(c, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Insurance created successfully", "Insurance": "ins"})
}
