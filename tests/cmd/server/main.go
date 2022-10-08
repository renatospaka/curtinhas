package main

import (
	"database/sql"
	"net/http"

	"github.com/renatospaka/tests/internal/infra/controller"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	controllers := controller.NewBaseHandler(db)
	http.HandleFunc("/clients", controllers.CreateClientHandler)
	http.ListenAndServe(":3055", nil)
}