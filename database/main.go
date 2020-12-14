package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", stringConexao)
	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()

	if error = db.Ping(); error != nil {
		log.Fatal(error)
	}

	fmt.Println("Conexão aberta com o banco de dados")

	resultSet, error := db.Query("select * from usuarios")
	if error != nil {
		log.Fatal(error)
	}

	defer resultSet.Close()


}
