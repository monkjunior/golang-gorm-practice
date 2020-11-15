package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dns := "root:admin@tcp(127.0.0.1:3306)/gorm"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	_ = db
}
