package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.StrictSlash(true)
	r.Handle("/", IO(HomeHandler)).Methods(http.MethodGet)
}
