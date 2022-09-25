package config

import (
	"database/sql"
	"fmt"

	_ "gorm.io/driver/mysql"
)

const (
	host           = "localhost"
	port           = "3306"
	user           = "root"
	password       = "rroooott"
	dbname         = "orders_by"
	dbMaxIdle      = 4
	dbMaxOpenConns = 25
)

// cara singleton
// var (
// 	DB *sql.DB
// )

// func ConnectPostgres() error {
// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		return err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}

// 	DB = db

// 	return nil
// }

// func GetDB() *sql.DB {
// 	return DB
// }

func ConnectMySQL() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user,password, host, port, dbname)
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}


	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxOpenConns)
	db.SetMaxIdleConns(dbMaxIdle)

	return db, nil
}
