package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/u-shylianok/golang-test-tasks/task05/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	ads := router.Group("/ads")
	{
		ads.POST("/", h.createAd)
		ads.GET("/", h.listAds)
		ads.GET("/:id", h.getAd)
	}

	return router
}
