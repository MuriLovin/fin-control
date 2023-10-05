package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Inst  *sql.DB
	Error error
}

func InitDB() Database {
	db, err := sql.Open("mysql", "developer:dev@(172.19.0.2:3306)/myfinance")
	var instance Database = Database{db, err}
	return instance
}
