package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Post struct {
	ID      uint
	Title   string
	Author  string
	Content string
}

func main() {
	dns := "root:admin@tcp(127.0.0.1:3306)/gorm"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to mysql db: %+v", err)
	}

	err = db.AutoMigrate(User{})
	if err != nil {
		log.Fatalf("Could not auto migrate User table to mysql db: %+v", err)
	}

	err = db.AutoMigrate(Post{})
	if err != nil {
		log.Fatalf("Could not auto migrate Post table to mysql db: %+v", err)
	}
}
