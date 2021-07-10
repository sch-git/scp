package models

type SCP struct {
	ID   int64  `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}

func (SCP) TableName() string {
	return "scp"
}
