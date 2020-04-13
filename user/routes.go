package user

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func listUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "listing users...")
}

func UserRouter() chi.Router {
	// func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", listUsers)
	return r
}
