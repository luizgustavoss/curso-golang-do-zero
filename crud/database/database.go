package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // database driver
)

func Connect() (*sql.DB, error) {

	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", stringConexao)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil
}
