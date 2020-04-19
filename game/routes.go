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

func createGame(playerId int, dif models.Difficuty, w http.ResponseWriter) models.Game {
	fmt.Println("creating game...", playerId, dif)
	if !models.IsPlayerValid(playerId) {
		fmt.Fprintf(w, "Player of id %d not in the database.", playerId)
	}
	game := models.NewGame(playerId, dif)
	return game
}

func startGame(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Id  int
		Dif models.Difficuty
	}
	json.NewDecoder(r.Body).Decode(&data)
	newGame := createGame(data.Id, data.Dif, w)
	json.NewEncoder(w).Encode(newGame.Id)
}

func GameRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", listGames)
	r.Get("/{id}", findGame)
	r.Post("/", startGame)

	return r
}
