package main

import (
	"flag"
	"log"

	"github.com/codescalersinternships/todoapp-omar/internal"
)

func main() {
	var dbFilePath string
	flag.StringVar(&dbFilePath, "d", "./todoapp.db", "Specify the filepath of sqlite database")
	flag.Parse()

	// database client declaration
	client := internal.DBClient{}

	app := internal.NewApp(client)

	if err := app.Run(dbFilePath); err != nil {
		log.Fatal(err)
	}
}
