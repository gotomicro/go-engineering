package mysql

type Resource struct {
	Id       int64  `gorm:"not null"`
	Title    string `gorm:"not null;type:varchar(160)"`
	Nickname string `gorm:"not null;type:varchar(160)"`
	Content  string `gorm:"not null;type:longtext"`
}

func (Resource) TableName() string {
	return "resource"
}
