package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/db-name")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Query("Insert into products values('ABC', 'Carrinho')")
}
