package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Age       uint
	Biarthday time.Time
}

func main() {
	db, err := RDBConnect()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})
}

type DB struct {
	*gorm.DB
}

func RDBConnect() (*gorm.DB, error) {
	db, err := createDB()
	if err != nil {
		return nil, err
	}
	d := &DB{db}
	return d, nil
}

func createDB() (*gorm.DB, error) {
	dsn := "host=" + "sample" + " user=" + "sample" + " password=" + "sample" + " port=" + "5432" + " sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to db", err)
		return nil, err
	}

	return db, nil
}
