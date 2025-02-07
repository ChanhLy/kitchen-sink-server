package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"go-server/utils"
	"log"
	"os"
	"strings"
)

var db *Queries
var connection *sql.DB

// GetDb Singleton DB Connection for project
func GetDb() *Queries {
	if db != nil {
		return db
	}

	connection = dbConnection()

	cfg := utils.GetConfig()
	if cfg.DB.Path == ":memory:" {
		execSqlFile(connection, "/db/schema.sql")
	}

	db = New(connection)
	return db
}

func CloseDb() {
	if connection == nil {
		log.Println("Database connection is closed")
		return
	}

	err := connection.Close()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Database connection closed")
}

func dbConnection() *sql.DB {
	log.Println("Connecting to database...")
	cfg := utils.GetConfig()

	connection, openDbErr := sql.Open("sqlite3", cfg.DB.Path)
	if openDbErr != nil {
		log.Fatal(openDbErr)
	}

	// Confirm a successful connection.
	if err := connection.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
	return connection
}

func closeConnection(connection *sql.DB) {
	if err := connection.Close(); err != nil {
		log.Fatal(err)
	}
}

func execSqlFile(connection *sql.DB, path string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cfg := utils.GetConfig()

	// Tests are run with test file's working directory, not from kitchen-sink-server directory
	// Ex: .../kitchen-sink-server/users/user_test.go will have base wd equal .../kitchen-sink-server/users, not .../kitchen-sink-server
	sqlFilePath := strings.Split(wd, cfg.ProjectPath)[0] + cfg.ProjectPath + path

	file, err := os.ReadFile(sqlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Execute all
	log.Println("Execute " + path)
	_, err = connection.Exec(string(file))
	if err != nil {
		log.Fatal(err)
	}
}
