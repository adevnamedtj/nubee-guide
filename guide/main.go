package main

import (
	"log"
	"net/http"

	"github.com/ckalagara/nubee-guide/guide/contentmetadata"
)

func main() {
	database, err := contentmetadata.CreateDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database: %s", err.Error())
	}

	contentService := &contentService.ContentService{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	contentService.setupRouter()

	log.Fatal(http.ListenAndServe(":8080", contentService.Router))
}
