package repositories

import (
	"mock-gm-home-service/structures"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TbCustomerRepo interface {
	GetCustomers(c *gin.Context, status string) ([]structures.Member, error)
}

func NewTbTCustomerRepo(db *gorm.DB) TbCustomerRepo {
	return &tbTCustomerRepo{
		db: db,
	}
}

type tbTCustomerRepo struct {
	db *gorm.DB
}

// GetCustomers implements TbTCustomerRepo.
func (find *tbTCustomerRepo) GetCustomers(c *gin.Context, status string) ([]structures.Member, error) {
	query := find.db.Model(&structures.Member{})
	resp := []structures.Member{}
	res := query.Where("status = ?", status).
		Find(&resp)
	if res.Error != nil {
		return nil, res.Error
	}
	return resp, nil
}
