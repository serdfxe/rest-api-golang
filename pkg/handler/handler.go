package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/serdfxe/rest-api-golang/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		admin := api.Group("/admin")
		{
			admin.GET("/all", h.getAll)
		}
	}

	return router
}
