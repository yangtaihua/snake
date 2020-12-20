package seeders

import "github.com/jinzhu/gorm"

//数据写入
func Seeder(db gorm.DB) {
	SysTreeSeeder(db)
}
