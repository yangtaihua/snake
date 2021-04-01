package seeders

import "github.com/1024casts/snake/internal/model"

func SysTreeSeeder() {
	db := model.GetDB()
	db.AutoMigrate()
}
