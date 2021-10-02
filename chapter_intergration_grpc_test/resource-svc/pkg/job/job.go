package job

import (
	"fmt"

	"github.com/gotomicro/ego/task/ejob"
	"go-engineering/resource-svc/pkg/invoker"
	"go-engineering/resource-svc/pkg/model/mysql"
)

func InstallComponent(ctx ejob.Context) error {
	models := []interface{}{
		&mysql.Resource{},
	}
	gormdb := invoker.Db.Debug()
	err := gormdb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
	if err != nil {
		return err
	}
	fmt.Println("create table ok")
	return nil
}

func InitializeComponent(ctx ejob.Context) error {
	create := mysql.Resource{
		Id:       1,
		Title:    "测试文章",
		Nickname: "ego",
		Content:  "测试文章内容",
	}
	return invoker.Db.Create(&create).Error
}
