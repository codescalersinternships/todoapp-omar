package main

import (
	"flag"
	"log"

	_ "github.com/codescalersinternships/todoapp-omar/docs"

	"github.com/codescalersinternships/todoapp-omar/app"
)

// @title 			Todoapp API
// @version			1.0
// @description A Todoapp API in Go using Gin framework and sqlite3
// @host 				localhost:8080
func main() {
	var dbFilePath string
	flag.StringVar(&dbFilePath, "d", "./todoapp.db", "Specify the filepath of sqlite database")
	flag.Parse()

	app, err := app.NewApp(dbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(dbFilePath); err != nil {
		log.Fatal(err)
	}
}
