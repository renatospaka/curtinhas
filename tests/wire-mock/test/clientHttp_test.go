package test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/renatospaka/tests/internal/infra/controller"
)


func setupDb() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	sql := `
			CREATE TABLE IF NOT EXISTS clients (
				id TEXT NOT NULL PRIMARY KEY,
				name TEXT,
				email TEXT,
				points INTEGER
			);
		`
	_, err = db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return db
}

func tearDownDb(db *sql.DB) {
	db.Exec("DROP TABLE clients;")
	db.Close()
}


func TestCreateClientHandler(t *testing.T) {
	db := setupDb()
	defer tearDownDb(db)

	control := controller.NewBaseHandler(db)
	t.Run("Should create a client", func(t *testing.T) {
		data := `{"name": "John Doe", "email": "jdoe@notemail.com"}`
		reader := strings.NewReader(data)
		request, _ := http.NewRequest("POST", "/clients", reader)
		response := httptest.NewRecorder()

		control.CreateClientHandler(response, request)
		if response.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, response.Code)
		}
	})
}