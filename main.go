package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "gosql:20152008Gus!@tcp(127.0.0.1:3306)/ejemplodb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	//fmt.Println(db.Name())
	//fmt.Print(dsn)

	type User struct {
		ID    uint
		Name  string
		Age   int
		Fecha time.Time
		// hundreds of fields
	}

	user := User{Name: "Gonzalo", Age: 45, Fecha: time.Now()}
	result := db.Create(&user) // pass pointer of data to Create
	//user.ID             // returns inserted data's primary key
	//result.Error        // returns error
	//result.RowsAffected // returns inserted records count

	fmt.Print(result.RowsAffected)

}
