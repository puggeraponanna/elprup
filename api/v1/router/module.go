package router

import (
	"elprup/api/v1/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func InitModuleRouter(r *mux.Router) {
	sr := r.PathPrefix("/module").Subrouter()
	sr.HandleFunc("/{moduleId}/record/{recordId}", handler.RecordHandler).Methods(http.MethodGet)
	sr.HandleFunc("/{moduleId}/record", handler.RecordsHandler).Methods(http.MethodGet)
}
