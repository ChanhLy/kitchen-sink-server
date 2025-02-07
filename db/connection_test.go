package database

import (
	"database/sql"
	"go-server/utils"
	"log"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestCreateDbSchemas(t *testing.T) {
	type args struct {
		connection *sql.DB
		path       string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Create Schemas",
			args: args{
				path: "/db/schema.sql",
			},
		},
	}
	connection := dbConnection()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			execSqlFile(connection, tt.args.path)
		})
	}
	closeConnection(connection)
}

func TestMigrateDbSchemas(t *testing.T) {
	migrationUpFilesPaths, migrationDownFilesPaths := getMigrationFilesPath()
	connection := dbConnection()

	type args struct {
		connection *sql.DB
		paths      []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Migrate Up Schemas",
			args: args{
				paths: migrationUpFilesPaths,
			},
		},
		{
			name: "Migrate Down Schemas",
			args: args{
				paths: migrationDownFilesPaths,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, path := range tt.args.paths {
				execSqlFile(connection, path)
			}
		})
	}
	closeConnection(connection)
}

func getMigrationFilesPath() ([]string, []string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cfg := utils.GetConfig()
	migrationFolderPath := strings.Split(wd, cfg.ProjectPath)[0] + cfg.ProjectPath + "/db/migrations"
	files, err := os.ReadDir(migrationFolderPath)
	if err != nil {
		log.Fatal(err)
	}

	var migrationUpFiles []string
	var migrationDownFiles []string

	for _, file := range files {
		if strings.Contains(file.Name(), ".up.sql") {
			migrationUpFiles = append(migrationUpFiles, "/db/migrations/"+file.Name())
		}
		if strings.Contains(file.Name(), ".down.sql") {
			migrationDownFiles = append(migrationDownFiles, "/db/migrations/"+file.Name())
		}
	}
	slices.Reverse(migrationDownFiles)
	return migrationUpFiles, migrationDownFiles
}
