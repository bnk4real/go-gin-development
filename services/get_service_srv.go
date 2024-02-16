package services

import (
	"github.com/gin-gonic/gin"
)

type ServiceRequest struct {
	ServiceId string `json:"serviceid"`
}

type ServiceResponse struct {
	ServiceId    string       `json:"serviceid"`
	ServiceName  string       `json:"servicename"`
	ServiceDesc  string       `json:"servicedesc"`
	ResponseData ResponseData `json:"responseData"`
}

type ServiceResponseData struct {
	ServiceId   string `json:"serviceid"`
	ServiceName string `json:"servicename"`
	ServiceDesc string `json:"servicedesc"`
}

func (srv *mainService) GetService(ctx *gin.Context, req *ServiceRequest) (res *[]ServiceResponse, e error) {

	findService, e := srv.tbService.GetService(ctx, req.ServiceId)
	if e != nil {
		return nil, e
	}
	if len(findService) <= 0 {
		return nil, e
	}

	result := []ServiceResponse{}

	for _, v := range findService {
		data := ServiceResponse{
			ServiceId:   v.ServiceId,
			ServiceName: v.ServiceName,
			ServiceDesc: v.ServiceDesc,
		}
		result = append(result, data)
	}

	return &result, nil
}
