package contentmetadata

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)
 // Content metedata crud operations
func (databasePointer *Database) findContentMetadataByID(id string) &Content, error {
	content := &Content{}
	err := databasePointer.QueryRow("SELECT video_id, title, description FROM `contentmetadata` WHERE video_id = ?", id).Scan(&Content.VideoID, &Content.Title, &Content.Description)
	if err != nil {
		log.Fatal("Failed to pulled data", err)
	}
	log.Println("Fetched data for content ", id)

	return content, err
}

func (databasePointer *Database) saveContentMetadata(content Content) error { 
	_, err := databasePointer.Exec("INSERT INTO `test` ----------etc....")
	if err != nil {
		log.Fatal("Database to insert content ", content)
	}
	return err
}

// database base functions
func createDatabase() (*sql.DB, error) {
	serverName := "localhost:3306"
	user := "myuser"
	password := "pw"
	dbName := "demo"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err := migrateDatabase(db); err != nil {
		return db, err
	}

	return db, nil
}

func migrateDatabase(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/db/migrations", dir),
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	migration.Log = &MigrationLogger{}

	migration.Log.Printf("Applying database migrations")
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	version, _, err := migration.Version()
	if err != nil {
		return err
	}

	migration.Log.Printf("Active database version: %d", version)

	return nil
}


