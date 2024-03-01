package routes

import (
	"heitortanoue/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Não foi possível parsear o ID do evento"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível recuperar o evento"})
		return
	}

	userId := c.GetInt64("userId")
	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível registrar o usuário para o evento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário registrado para o evento com sucesso"})
}

func unregisterForEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Não foi possível parsear o ID do evento"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível recuperar o evento"})
		return
	}

	userId := c.GetInt64("userId")
	err = event.Unregister(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível desregistrar o usuário para o evento: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário desregistrado para o evento com sucesso"})
}

func getEventRegistrations(c *gin.Context) {
	userId := c.GetInt64("userId")
	events, err := models.GetAllRegistrationsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível recuperar os eventos registrados"})
		return
	}

	c.JSON(http.StatusOK, events)
}
