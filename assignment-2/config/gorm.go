package config

import (
	"assignment2/httpserver/repositories/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func ConnectMySQLGORM() (*gorm.DB, error) {
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user,password, host, port, dbname)

	db, err := gorm.Open("mysql", dsn)

	db.AutoMigrate(&models.Order{}, &models.Item{})
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	// db.Debug().AutoMigrate(models.Order{})

	return db, nil
}
