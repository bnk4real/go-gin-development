package controllers

import (
	"mock-gm-home-service/services"

	"github.com/gin-gonic/gin"
)

type MainControllers interface {
	GetService(c *gin.Context)
	GetCustomer(c *gin.Context)
}

type mainController struct {
	srv services.MainService
}

func NewCustomerController(
	srv services.MainService,
) MainControllers {
	return &mainController{
		srv,
	}
}
