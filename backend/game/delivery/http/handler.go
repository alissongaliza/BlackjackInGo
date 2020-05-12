package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alissongaliza/BlackjackInGo/backend/game"
	"github.com/alissongaliza/BlackjackInGo/backend/user"
	"github.com/alissongaliza/BlackjackInGo/utils"
	"github.com/go-chi/chi"
)

type GameHandler struct {
	gameUsecase game.UseCase
	userUsecase user.UseCase
	userRepo    user.Repository
	gameRepo    game.Repository
}

func NewGameHandler(r chi.Router, gameUsecase game.UseCase,
	userUsecase user.UseCase, userRepo user.Repository, gameRepo game.Repository) {
	handler := &GameHandler{gameUsecase: gameUsecase, userUsecase: userUsecase,
		userRepo: userRepo, gameRepo: gameRepo}
	r.Route("/games", func(r chi.Router) {
		r.Get("/", handler.listGames)
		r.Post("/", handler.startGame)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.findGame)
			r.Put("/", handler.continueGame)
			// 	r.Get("/", handler.findGame)
			r.Post("/play", handler.play)
		})
	})
}

func (gh *GameHandler) listGames(w http.ResponseWriter, r *http.Request) {
	// 0 means dont filter by user
	games := gh.gameRepo.ListGame(0)
	json.NewEncoder(w).Encode(games)
}

func (gh *GameHandler) findGame(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "finding games...")
}

func (gh *GameHandler) startGame(w http.ResponseWriter, r *http.Request) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		// fmt.Fprint(w, r)
	// 		fmt.Println(r)
	// 	}
	// }()
	var request utils.GameCreateRequest
	json.NewDecoder(r.Body).Decode(&request)
	if !gh.userUsecase.IsUserValid(request.UserId) {
		panic(fmt.Sprintf("User of id %d not in the database.", request.UserId))
	}
	userIsValid := gh.userUsecase.IsUserValid(request.UserId)
	if user := gh.userRepo.GetUser(request.UserId); userIsValid {
		game := gh.gameUsecase.CreateGame(user, request.Dif, request.Bet)
		game = gh.gameUsecase.StartNewGame(game)
		json.NewEncoder(w).Encode(game)
	}
}

func (gh *GameHandler) continueGame(w http.ResponseWriter, r *http.Request) {
	gameId := utils.StringToInt(chi.URLParam(r, "id"))
	game := gh.gameUsecase.GetGame(gameId)
	game = gh.gameUsecase.ContinueGame(game)

	json.NewEncoder(w).Encode(game)
}

func (gh *GameHandler) play(w http.ResponseWriter, r *http.Request) {
	gameId := utils.StringToInt(chi.URLParam(r, "id"))
	var body utils.UserActionRequest
	json.NewDecoder(r.Body).Decode(&body)
	var game models.Game
	switch body.Action {
	case utils.Hit:
		game = gh.userUsecase.Hit(gameId, true)
	case utils.Stand:
		game = gh.userUsecase.Stand(gameId)
	case utils.DoubleDown:
		game = gh.userUsecase.DoubleDown(gameId)
	}
	json.NewEncoder(w).Encode(game)

}
