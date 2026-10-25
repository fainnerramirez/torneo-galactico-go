package connection

import (
	"database/sql"
	"fmt"
)

func ConnectionFactory() *sql.DB {

	db, err := sql.Open("sqlite", "./torneo-db")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	fmt.Println("SQLite conection Exit!")

	return db
}
