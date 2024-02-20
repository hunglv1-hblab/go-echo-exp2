package db

import (
	"fmt"
	"fund-aplly-back/config"
	"fund-aplly-back/db/seeders"
	"fund-aplly-back/models/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name)

	fmt.Println(dataSourceName)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	migrations.Migration(db)

	userSeeder := seeders.NewUserSeeder(db)
	userSeeder.SetUsers()

	return db
}
