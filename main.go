package main

import (
	"biblioV2/db"
	"fmt"
)

func main() {
	t, err := db.OpenDBConnection()
	if err != nil {
		panic(err.Error())
	}
	defer t.Close()

	fmt.Println("Connexion à la base de données MySQL réussie!")

}
