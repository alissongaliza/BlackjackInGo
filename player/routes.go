package player

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/alissongaliza/BlackjackInGo/repository"

	"github.com/go-chi/chi"
)

type playerRequest struct {
	*models.Player
}

func listPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "listing players...")
}

func findPlayer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "finding players...")
}

func addPlayer(w http.ResponseWriter, r *http.Request) {
	var data models.Player
	json.NewDecoder(r.Body).Decode(&data)

	newPlayer := models.NewPlayer(data.Name, data.Age)
	json.NewEncoder(w).Encode(newPlayer)
	// fmt.Fprintf(w, "adding player...")
	// fmt.Fprintf(w, "%+v", newPlayer)

}

func hit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit called")
	var id int
	json.NewDecoder(r.Body).Decode(&id)

	// json.NewEncoder(w).Encode(newPlayer)

}

func stand(w http.ResponseWriter, r *http.Request) {
	fmt.Println("stand called")
	var data models.Player
	json.NewDecoder(r.Body).Decode(&data)

	newPlayer := models.NewPlayer(data.Name, data.Age)
	json.NewEncoder(w).Encode(newPlayer)

}

func doubleDown(w http.ResponseWriter, r *http.Request) {
	fmt.Println("doubleDown called")
	var data models.Player
	json.NewDecoder(r.Body).Decode(&data)

	newPlayer := models.NewPlayer(data.Name, data.Age)
	json.NewEncoder(w).Encode(newPlayer)

}

func PlayerRouter() chi.Router {
	// func PlayerRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", listPlayers)
	r.Post("/", addPlayer)
	r.Get("/{id}", findPlayer)
	r.Post("/{id}/hit", hit)
	r.Post("/{id}/stand", stand)
	r.Post("/{id}/doubleDown", doubleDown)

	return r
}
