package handler

import (
	"github.com/Bernar11296/rest-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
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
		lists := api.Group("/lists")
		{
			lists.POST("/", h.CreateList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.CreateItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
