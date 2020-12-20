package migrations

import (
	"github.com/1024casts/snake/internal/model"
	"github.com/jinzhu/gorm"
)

func Migrate(db gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&model.UserBaseModel{})
}
