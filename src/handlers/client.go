package handlers;

import (
	"net/http"
	"strconv"
	"travelagency/src/utils"
	"travelagency/prisma/db"
	"github.com/gin-gonic/gin"
)

func GetClients(c *gin.Context) {
	// var clients []db.InnerClient

	prismaClient := utils.GetPrisma(c)

	clients, err := prismaClient.Client.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clients})
}

func PostClient(c *gin.Context) {
	var payload db.InnerClient

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prismaClient := utils.GetPrisma(c)
	insertedClient, err := prismaClient.Client.CreateOne(
		db.Client.Name.Set(payload.Name),
		db.Client.BirthDate.Set(payload.BirthDate),
		db.Client.Email.Set(payload.Email),
		db.Client.Address.SetOptional(payload.Address),
		db.Client.Occupation.SetOptional(payload.Occupation),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client created successfully", "client": insertedClient})
}

func PatchClient(c *gin.Context) {
	var payload db.InnerClient

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prismaClient := utils.GetPrisma(c)
	updatedClient, err := prismaClient.Client.FindUnique(
		db.Client.ID.Equals(id),
	).Update(
		db.Client.Name.Set(payload.Name),
		db.Client.BirthDate.Set(payload.BirthDate),
		db.Client.Email.Set(payload.Email),
		db.Client.Address.SetOptional(payload.Address),
		db.Client.Occupation.SetOptional(payload.Occupation),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client patched", "client": updatedClient})
}

func DeleteClient(c *gin.Context) {
	// TODO: utilizar o flag_removed ao invés de apagar o registro na tabela

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	prismaClient := utils.GetPrisma(c)
	deletedClient, err := prismaClient.Client.FindUnique(
		db.Client.ID.Equals(id),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully", "id": deletedClient.ID})
}
