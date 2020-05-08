package game

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/alissongaliza/BlackjackInGo/backend/repository"
	"github.com/alissongaliza/BlackjackInGo/utils"
	"github.com/go-chi/chi"
)

func listGames(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.URL.Query()["userId"]
	if !ok {
		userId[0] = ""
	}
	games := models.GetGameDb().List(userId[0])
	fmt.Println("games found", len(games))
	json.NewEncoder(w).Encode(games)
}

func findGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "finding games...")
}

func createGame(userId int, dif utils.Difficulty, bet int, w http.ResponseWriter) models.Game {
	fmt.Println("creating game...", userId, dif, bet)
	if !models.IsUserValid(userId) {
		panic(fmt.Sprintf("User of id %d not in the database.", userId))
	}
	fmt.Println("User is valid")
	game := models.NewGame(userId, dif, bet)
	return models.GetGameDb().Create(game)
}

func startGame(w http.ResponseWriter, r *http.Request) {
	var data struct {
		UserId int
		Dif    utils.Difficulty
		Bet    int
	}
	defer func() {
		if r := recover(); r != nil {
			// fmt.Fprint(w, r)
			fmt.Println(r)
		}
	}()
	json.NewDecoder(r.Body).Decode(&data)
	newGame := createGame(data.UserId, data.Dif, data.Bet, w)

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
