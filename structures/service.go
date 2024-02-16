package structures

type Service struct {
	ServiceId     string `gorm:"column:serviceid"`
	ServiceTypeId string `gorm:"column:servicetypeid"`
	ServiceName   string `gorm:"column:servicename"`
	ServiceDesc   string `gorm:"column:servicedesc"`
	CusId         string `gorm:"column:cusid"`
}

func (*Service) TableName() string {
	return "tb_service"
}
