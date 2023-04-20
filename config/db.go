package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/tes-go")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("successfully connect database")
	return db
}
