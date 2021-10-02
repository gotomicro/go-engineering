package mysql

import (
	resourcev1 "go-engineering/proto/pb/resource/v1"
)

type Resource struct {
	Id       int64  `gorm:"not null"`
	Title    string `gorm:"not null;type:varchar(160)"`
	Nickname string `gorm:"not null;type:varchar(160)"`
	Content  string `gorm:"not null;type:longtext"`
}

type Resources []Resource

func (Resource) TableName() string {
	return "resource"
}

func (r Resource) ToDetailPb() (output *resourcev1.DetailResponse) {
	output = &resourcev1.DetailResponse{
		Title:    r.Title,
		Nickname: r.Nickname,
		Content:  r.Content,
	}
	return
}

func (rs Resources) ToListPb() (output []*resourcev1.Info) {
	for _, value := range rs {
		output = append(output, &resourcev1.Info{
			Id:       value.Id,
			Title:    value.Title,
			Nickname: value.Nickname,
		})
	}
	return
}
