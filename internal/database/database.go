package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Inst  *sql.DB
	Error error
}

var databaseInstance Database

func GetInstance() Database {
	databaseInstance_ptr := &databaseInstance
	if databaseInstance_ptr.Inst != nil {
		return databaseInstance
	}

	db, err := sql.Open(os.Getenv("DATABASE_DRIVER"), os.Getenv("DATABASE_URL"))
	*databaseInstance_ptr = Database{db, err}
	return databaseInstance
}
