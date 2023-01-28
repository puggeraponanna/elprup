package router

import (
	"elprup/api/v1/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.IndexHandler).Methods(http.MethodGet)
	InitModuleRouter(r)
	return r
}
