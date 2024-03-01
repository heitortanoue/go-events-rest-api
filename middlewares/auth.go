package middlewares

import (
	"heitortanoue/rest-api/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"message": "Você precisa estar logado para acessar este recurso"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "Token inválido: " + err.Error()})
		return
	}

	c.Set("userId", userId) // adiciona o userId ao contexto da requisição

	c.Next() // continua a execução
}
