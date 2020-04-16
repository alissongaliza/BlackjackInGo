package game

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/alissongaliza/BlackjackInGo/repository"
	"github.com/go-chi/chi"
)

func listGames(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "listing games...")
}

func findGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "finding games...")
}

func createGame(userId int, w http.ResponseWriter) models.Game {
	fmt.Println("creating game...")
	if !models.IsUserValid(userId) {
		fmt.Fprintf(w, "User of id %d not in the database.", userId)
	}
	var game models.Game
	return game
}

func startGame(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Id int
	}
	json.NewDecoder(r.Body).Decode(&data)
	createGame(data.Id, w)
}

func GameRouter() chi.Router {
	// func GameRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", listGames)
	r.Get("/{id}", findGame)
	r.Post("/", startGame)

	return r
}
