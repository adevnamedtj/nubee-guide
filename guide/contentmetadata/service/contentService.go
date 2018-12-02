package contentmetadata

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ContentService struct{
	Router *mux.Router
	Database *sql.DB
}

func (contentServicePointer *ContentService) setupRouter(){
	contentServicePointer.Router.Methods("GET").Path("/service/content/{id}").HandlerFunc(contentServicePointer.getContent)

	contentServicePointer.Router.Methods("PUT").Path("/service/content").HandlerFunc(contentServicePointer.saveContent)

	contentServicePointer.Router.Methods("POST").Path("/service/content").HandlerFunc(contentServicePointer.updateContent)

	contentServicePointer.Router.Methods("DELETE").Path("/service/content/{id}").HandlerFunc(contentServicePointer.deleteContent)
}


func (contentServicePointer *ContentService) getContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("Invalid/No request permater found, ID")
	}
	contentmeta, err := contentServicePointer.Database.findContentMetadataByID(id)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(contentmeta); err != nil {
		panic(err)
	}


func (contentServicePointer *ContentService) saveContent(w http.ResponseWriter, r *http.Request) {
}

func (contentServicePointer *ContentService) updateContent(w http.ResponseWriter, r *http.Request) {
}

func (contentServicePointer *ContentService) deleteContent(w http.ResponseWriter, r *http.Request) {
}