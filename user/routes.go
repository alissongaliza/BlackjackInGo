package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/alissongaliza/BlackjackInGo/repository"

	"github.com/go-chi/chi"
)

type userRequest struct {
	*models.User
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "listing users...")
}

func findUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "finding users...")
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var data models.User
	json.NewDecoder(r.Body).Decode(&data)

	newUser := models.NewUser(data.Name, data.Age)
	json.NewEncoder(w).Encode(newUser)
	// fmt.Fprintf(w, "adding user...")
	// fmt.Fprintf(w, "%+v", newUser)

}

func UserRouter() chi.Router {
	// func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", listUsers)
	r.Get("/{id}", findUser)
	r.Post("/", addUser)

	return r
}
