package services

import (
	"mock-gm-home-service/repositories"

	"github.com/gin-gonic/gin"
)

type MainService interface {
	GetService(*gin.Context, *ServiceRequest) (*[]ServiceResponse, error)
	GetCustomer(*gin.Context, *CustomerRequest) (*[]CustomerResponse, error)
}

type mainService struct {
	srv       repositories.TbCustomerRepo
	tbService repositories.TbTServiceRepo
}

func NewCustomerService(
	srv repositories.TbCustomerRepo,
	tbService repositories.TbTServiceRepo,
) MainService {
	return &mainService{
		srv:       srv,
		tbService: tbService,
	}
}

type ResponseData struct {
	Code     string `json:"code"`
	Messsage string `json:"message"`
	Status   string `json:"status"`
}
