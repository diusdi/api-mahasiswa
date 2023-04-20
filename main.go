package main

import (
	"api-mahasiswa/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.ConnectDb()
	insert, err := db.Query("INSERT INTO user VALUES ( 'TEST' )")

	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
