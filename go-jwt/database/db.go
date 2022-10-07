package database

import (
	"fmt"
	"go-jwt/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host = "localhost"
	user = "root"
	password = "rroooott"
	dbPort = "3306"
	dbname = "simple-api"
	db *gorm.DB
	err error
)

func StartDB(){
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host,user,password,dbname,dbPort)
	dsn := config
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil{
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("sukses koneksi ke database :")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB{
	return db
}