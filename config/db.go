package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() (*sql.DB, error) {
	dbName := "api-mahasiswa"
	return sql.Open("mysql", "root@tcp(127.0.0.1:3306)/"+dbName+"?parseTime=true")
}
