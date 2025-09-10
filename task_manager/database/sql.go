package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB 

func Mysql() {
dsn := "root:@tcp(127.0.0.1:3306)/taskmanager?parseTime=true&loc=UTC"
	var err error
	DB, err = sql.Open("mysql",dsn)
	if err!=nil{
		log.Fatal("Error opening DB:",err)
	}

	err=DB.Ping()
	if err!=nil{
		log.Fatal("Error connecting to db:",err)
	}

	fmt.Println("successfully connected to mysql")
}