package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go-server/utils"
	"log"
	"os"
	"strings"
)

var queries *Queries

func ConnectDB() *Queries {
	log.Println("Connecting to database...")
	cfg, errConfig := utils.LoadConfig()
	if errConfig != nil {
		log.Fatal(errConfig)
	}

	connection, openDbErr := sql.Open("sqlite3", cfg.DB.Path)

	if openDbErr != nil {
		log.Fatal(openDbErr)
	}

	// Confirm a successful connection.
	if err := connection.Ping(); err != nil {
		log.Fatal(err)
	}

	if cfg.DB.Path == ":memory:" {
		runDbSchemas(connection)
	}

	return New(connection)
}

func runDbSchemas(connection *sql.DB) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Tests are run with test file's working directory, not from kitchen-sink-server directory
	// Ex: .../kitchen-sink-server/users/user_test.go will have base wd equal .../kitchen-sink-server/users, not .../kitchen-sink-server
	schemaFilePath := strings.Split(wd, "/kitchen-sink-server")[0] + "/kitchen-sink-server/db/schema.sql"

	file, err := os.ReadFile(schemaFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Execute all
	_, err = connection.Exec(string(file))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetQueries() *Queries {
	if queries == nil {
		queries = ConnectDB()
	}

	return queries
}
