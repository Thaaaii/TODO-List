package main

import (
	"github.com/Thaaaii/TODO-List/database"
	"github.com/Thaaaii/TODO-List/routes"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.InitDatabase()
	//database.CreateTables()
	routes.InitRouter()
}
