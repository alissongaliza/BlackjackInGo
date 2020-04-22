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

func createGame(playerId int, dif models.Difficuty, bet int, w http.ResponseWriter) models.Game {
	fmt.Println("creating game...", playerId, dif, bet)
	if !models.IsPlayerValid(playerId) {
		panic(fmt.Sprintf("Player of id %d not in the database.", playerId))
	}
	fmt.Println("Player is valid")
	game := models.NewGame(playerId, dif, bet)
	return game
}

func startGame(w http.ResponseWriter, r *http.Request) {
	var data struct {
		PlayerId int
		Dif      models.Difficuty
		Bet      int
	}
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		// fmt.Fprint(w, r)
	// 		fmt.Println(r)
	// 	}
	// }()
	json.NewDecoder(r.Body).Decode(&data)
	newGame := createGame(data.PlayerId, data.Dif, data.Bet, w)
	fmt.Println("Game Created")
	game := models.StartGame(newGame.Id)
	json.NewEncoder(w).Encode(game)
}

func GameRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", listGames)
	r.Get("/{id}", findGame)
	r.Post("/", startGame)

	return r
}
