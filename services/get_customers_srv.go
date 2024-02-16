package services

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	StatusActive = "active"
)

type CustomerRequest struct {
	Status string `json:"status"`
}

type CustomerResponse struct {
	CusId  string `json:"cusId"`
	Status string `json:"status"`
}

func (srv *mainService) GetCustomer(ctx *gin.Context, req *CustomerRequest) (res *[]CustomerResponse, err error) {

	// validate req
	if req.Status != "active" {
		err := errors.New("data not found")
		return nil, err
	}

	resp, err := srv.srv.GetCustomers(ctx, req.Status)
	if err != nil {
		return nil, err
	}

	result := []CustomerResponse{}

	if len(resp) > 0 {
		for _, v := range resp {
			datas := CustomerResponse{
				CusId:  v.CusId,
				Status: v.Status,
			}
			result = append(result, datas)
		}
	}

	return &result, nil
}
