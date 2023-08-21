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
	var port int
	flag.StringVar(&dbFilePath, "d", "./todoapp.db", "Specify the filepath of sqlite database")
	flag.IntVar(&port, "p", 8080, "Specify the port number")
	flag.Parse()

	app, err := app.NewApp(dbFilePath, port)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
