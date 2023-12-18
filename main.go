package main

import (
	"github.com/Thaaaii/TODO-List/backend"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	backend.InitServer()
}
