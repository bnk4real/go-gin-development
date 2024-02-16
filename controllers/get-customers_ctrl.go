package controllers

import (
	"mock-gm-home-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *mainController) GetCustomer(ctx *gin.Context) {
	req := services.CustomerRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation Error"})
		return
	}

	result, err := c.srv.GetCustomer(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
