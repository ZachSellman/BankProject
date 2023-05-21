package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// temporarily defined as constants, in future I will use env variables instead.
const (
	// dbDriver requires an imported database driver to talk to a specific db engine (in our case lib/pq)
	dbDriver = "postgres"
	// This is copied from Make file migrate command
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

// testQueries defined as Global variable because we will use it extensively in all of our unit tests
var testQueries *Queries

// Opens a connection to the database and assigns the value of testQueries to the connection object
// By convention, the TestMain function is used as the entry point of all unit test functions inside a specific test package
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
