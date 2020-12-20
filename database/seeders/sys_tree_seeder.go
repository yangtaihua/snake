package seeders

import "github.com/jinzhu/gorm"

func SysTreeSeeder(db gorm.DB) {
	db.AutoMigrate()
}
