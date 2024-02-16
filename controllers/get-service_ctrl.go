package controllers

import (
	"mock-gm-home-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *mainController) GetService(ctx *gin.Context) {
	req := services.ServiceRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation Error"})
		return
	}
	result, err := c.srv.GetService(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Internal Server Error"})
	}
	ctx.JSON(http.StatusOK, result)
}
