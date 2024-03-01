package routes

import (
	"heitortanoue/rest-api/models"
	"heitortanoue/rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados estão no formato errado, tente novamente: " + err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível salvar o usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário criado com sucesso"})
}

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados estão no formato errado, tente novamente: " + err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Não foi possível autenticar o usuário: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login feito com sucesso", "token": token})
}
