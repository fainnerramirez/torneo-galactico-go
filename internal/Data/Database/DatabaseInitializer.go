package database

import (
	connection "api-torneo/internal/Data/Connection"
	"fmt"
)

func Initialize() {

	db, err := connection.ConnectionFactory()

	if err != nil {
		fmt.Println(err)
		return
	}
}
