package main

import (
	"github.com/Thaaaii/TODO-List/backend"
	"github.com/Thaaaii/TODO-List/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	models.InitDatabase()
	backend.InitServer()
}
