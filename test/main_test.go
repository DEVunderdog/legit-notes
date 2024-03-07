package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	database "github.com/DEVunderdog/legit-notes/database/sqlc"
	"github.com/DEVunderdog/legit-notes/utils"
	_ "github.com/lib/pq" // The idea for the sake of knowledge it just distilled
)

var testQueries *database.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../")

	if err != nil {
		log.Fatal("We get config err: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = database.New(testDB)

	os.Exit(m.Run())
	
}