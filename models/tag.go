package models

type Tag struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);unique;not null" json:"name"`
}

func (Tag) TableName() string {
	return "tags"
}
