package database_test

import (
	"database/sql"
	"testing"
	"github.com/stretchr/testify/assert"
	_ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"

	"github.com/renatospaka/tests/internal/entity"
	"github.com/renatospaka/tests/internal/infra/database"
)

func TestClientRepository_Save(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.Nil(t, err)
	defer db.Close()

	sql := `
			CREATE TABLE IF NOT EXISTS clients (
				id TEXT NOT NULL PRIMARY KEY,
				name TEXT,
				email TEXT,
				points INTEGER
			);
		`
	_, err = db.Exec(sql)
	assert.Nil(t, err)

	clientRepo := database.NewClientRepository(db)
	client := &entity.Client{
		Name:   "John Doe",
		Email:  "jdoe@notemail.com",
		ID:     "ID123",
		Points: 0,
	}
	err = clientRepo.Save(client)
	assert.Nil(t, err)

	var id, name, email string
	var points int
	err = db.QueryRow("SELECT id, name, email, points FROM clients WHERE id = $1", client.ID).Scan(
		&id, &name, &email, &points)
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", name)
	assert.Equal(t, "jdoe@notemail.com", email)
	assert.Equal(t, 0, points)
}

func BenchmarkClientRepositoryMemory_Save(b *testing.B) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.Nil(b, err)
	defer db.Close()

	sql := `
			CREATE TABLE IF NOT EXISTS clients (
				id TEXT NOT NULL PRIMARY KEY,
				name TEXT,
				email TEXT,
				points INTEGER
			);
		`
	_, err = db.Exec(sql)
	assert.Nil(b, err)

	clientRepo := database.NewClientRepository(db)

	for i := 0; i < b.N; i++ {
		client := &entity.Client{
			ID:     uuid.New().String(),
			Name:   "John Doe",
			Email:  "jdoe@notemail.com",
			Points: 0,
		}
		err = clientRepo.Save(client)
		assert.Nil(b, err)

		var id, name, email string
		var points int
		err = db.QueryRow("SELECT id, name, email, points FROM clients WHERE id = $1", client.ID).Scan(
			&id, &name, &email, &points)
		assert.Nil(b, err)
	}
}

func BenchmarkClientRepositoryDisk_Save(b *testing.B) {
	db, err := sql.Open("sqlite3", "test.db")
	assert.Nil(b, err)
	defer db.Close()

	sql := `
			CREATE TABLE IF NOT EXISTS clients (
				id TEXT NOT NULL PRIMARY KEY,
				name TEXT,
				email TEXT,
				points INTEGER
			);
		`
	_, err = db.Exec(sql)
	assert.Nil(b, err)

	clientRepo := database.NewClientRepository(db)

	for i := 0; i < b.N; i++ {
		client := &entity.Client{
			ID:     uuid.New().String(),
			Name:   "John Doe",
			Email:  "jdoe@notemail.com",
			Points: 0,
		}
		err = clientRepo.Save(client)
		assert.Nil(b, err)

		var id, name, email string
		var points int
		err = db.QueryRow("SELECT id, name, email, points FROM clients WHERE id = $1", client.ID).Scan(
			&id, &name, &email, &points)
		assert.Nil(b, err)
	}
}


// go test -v -bench=. (in the test files folder)
// go test -bench=. (in the test files folder)