package routes

import (
	"heitortanoue/rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// rotas que NÃO precisam de autenticação
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// rotas que PRECISAM de autenticação
	authenticated.POST("/events", createEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", unregisterForEvent)

	// rotas que PRECISAM de autenticação && só o dono do evento pode acessar
	authenticated.GET("/events/registrations", getEventRegistrations)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// rotas para autenticação
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
