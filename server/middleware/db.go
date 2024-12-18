package middleware

import (
	"github.com/hurricane5250/MyGoProject/server/log"
	"github.com/hurricane5250/MyGoProject/server/model"
)

// MakeMigrations 迁移数据库表
func MakeMigrations() {
	var err error
	err = model.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&model.User{},
		)
	if err != nil {
		log.Error("[ error ] 生成数据库表失败")
		return
	}
	log.Info("[ success ] 生成数据库表成功")
}
