package db

import (
	"database/sql"
	"log"
	"os"
	"simplebank/util"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

// Add "github.com/lib/pq" before execituing this test
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can not load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
