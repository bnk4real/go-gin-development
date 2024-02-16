package repositories

import (
	"mock-gm-home-service/structures"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TbTServiceRepo interface {
	GetService(c *gin.Context, serviceId string) ([]structures.Service, error)
}

// Provide Interface to use the functions within this interface
func NewTbTServiceRepo(db *gorm.DB) TbTServiceRepo {
	return &tbTServiceRepo{
		db: db,
	}
}

// return gorm.DB as custom struct
type tbTServiceRepo struct {
	db *gorm.DB
}

// GetCustomers implements TbTCustomerRepo.
func (find *tbTServiceRepo) GetService(c *gin.Context, serviceId string) ([]structures.Service, error) {
	query := find.db.Model(&structures.Service{})
	resp := []structures.Service{}
	res := query.Where("serviceId = ?", serviceId).
		Find(&resp)
	if res.Error != nil {
		return nil, res.Error
	}
	return resp, nil
}
