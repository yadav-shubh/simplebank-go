package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	// Set up a connection pool using pgxpool
	connString := "postgres://root:root@localhost:5432/simplebank?sslmode=disable"
	var err error
	testDB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("Postgres connected successfully")

	// Initialize testQueries using pgxpool connection
	testQueries = New(testDB)

	// Ensure the connection pool is closed after testing
	defer testDB.Close()

	// Run the tests
	os.Exit(m.Run())
}
