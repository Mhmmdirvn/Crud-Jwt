package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:Muhammadirvan011206@tcp(127.0.0.1:3306)/db_crud"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("faied connect to database")
	}

	db.AutoMigrate(&Person{})
	
	return db, err
}