package structures

type Member struct {
	CusId     string `gorm:"column:cusid"`
	Firstname string `gorm:"column:firstname"`
	Lastname  string `gorm:"column:lastname"`
	Tel       string `gorm:"column:mobileno"`
	Status    string `gorm:"column:status"`
}

func (*Member) TableName() string {
	return "tb_customers"
}
