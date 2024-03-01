package routes

import (
	"heitortanoue/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Aconteceu um erro ao retornar os eventos: " + err.Error()})
	}

	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Não foi possível parsear o ID"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível recuperar o evento"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	userId := c.GetInt64("userId")

	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Aconteceu um erro ao parsear o JSON: " + err.Error()})
		return
	}

	// Set the user ID
	event.UserID = userId

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Aconteceu um erro ao salvar no banco de dados: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Evento adicionado com sucesso", "event": event})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Não foi possível parsear o ID"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível recuperar o evento"})
		return
	}

	userId := c.GetInt64("userId")
	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"message": "Você não tem permissão para modificar este evento"})
		return
	}

	var updatedEvent models.Event // Não contem o ID nem o UserID
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Aconteceu um erro ao parsear o JSON: " + err.Error()})
		return
	}

	// Setar ID e UserId do evento
	updatedEvent.ID = id
	updatedEvent.UserID = userId

	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Aconteceu um erro ao atualizar: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Evento modificado com sucesso", "event": updatedEvent})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Não foi possível parsear o ID"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível recuperar o evento: " + err.Error()})
		return
	}

	userId := c.GetInt64("userId")
	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"message": "Você não tem permissão para deletar este evento"})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível deletar o evento: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Evento deletado com sucesso"})
}
