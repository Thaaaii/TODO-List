package main

import (
	"github.com/Thaaaii/TODO-List/backend"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//backend.CreateTables()
	//
	//lastUserID, _ := backend.InsertUserIntoTable("user1", "passwort")
	//lastTaskID, _ := backend.InsertTaskIntoTable("Sachen kaufen", "Weihnachtsgeschenke für Schwester", false, lastUserID)
	//backend.InsertCategoryIntoTable("Freizeit", lastTaskID)
	//
	//lastUserID, _ = backend.InsertUserIntoTable("user2", "admin")
	//lastTaskID, _ = backend.InsertTaskIntoTable("TODO-Listen App", "Frontend erstellen", true, lastUserID)
	//_, _ = backend.InsertCategoryIntoTable("Studium", lastTaskID)

	backend.InitServer()
}
