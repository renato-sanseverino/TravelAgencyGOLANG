package handlers

import (
	"net/http"
	"travelagency/prisma/db"
	"travelagency/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		db.Client.ID.Set(payload.ID),
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

	prismaClient := utils.GetPrisma(c)
	updatedClient, err := prismaClient.Client.FindUnique(
		db.Client.ID.Equals(id.String()),
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
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	prismaClient := utils.GetPrisma(c)
	deletedClient, err := prismaClient.Client.FindUnique(
		db.Client.ID.Equals(id.String()),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully", "id": deletedClient.ID})
}
