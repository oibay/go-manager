package connection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func PingDB(db *sql.DB) error {
	err := db.Ping()
	return err
}

func Initialize() *sql.DB {
	db, err := sql.Open("mysql", "mysql:123@tcp(db)/lms")
	if err != nil {
		panic(err)
	}
	return db
}