package migrations

import (
	"github.com/1024casts/snake/internal/model"
)

func Migrate() {
	db := model.GetDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&model.UserBaseModel{})
}
