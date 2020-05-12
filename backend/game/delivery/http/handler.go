package http

import (
	"encoding/json"
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
	if user, err := gh.userRepo.GetUser(request.UserId); err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		isValid := gh.userUsecase.IsUserValid(user)
		if !isValid {
			json.NewEncoder(w).Encode("User not Valid")
		}
		if game, err2 := gh.gameUsecase.CreateGame(user, request.Dif, request.Bet); err2 != nil {
			json.NewEncoder(w).Encode(err2)
		} else {
			var err3 error
			if game, err3 = gh.gameUsecase.StartNewGame(game); err3 != nil {
				json.NewEncoder(w).Encode(err3)
			} else {
				json.NewEncoder(w).Encode(game)
			}
		}
	}

}

func (gh *GameHandler) continueGame(w http.ResponseWriter, r *http.Request) {
	gameId := utils.StringToInt(chi.URLParam(r, "id"))
	game, err := gh.gameRepo.GetGame(gameId)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	game, err2 := gh.gameUsecase.ContinueGame(game)
	if err2 != nil {
		json.NewEncoder(w).Encode(err2)
	}
	json.NewEncoder(w).Encode(game)

}

func (gh *GameHandler) play(w http.ResponseWriter, r *http.Request) {
	gameId := utils.StringToInt(chi.URLParam(r, "id"))
	var body utils.UserActionRequest
	json.NewDecoder(r.Body).Decode(&body)
	if game, err := gh.gameRepo.GetGame(gameId); err != nil {
		json.NewEncoder(w).Encode(err)
	} else {

		var err error
		switch body.Action {
		case utils.Hit:
			game, err = gh.userUsecase.Hit(game, true)
		case utils.Stand:
			game = gh.userUsecase.Stand(game)
		case utils.DoubleDown:
			game, err = gh.userUsecase.DoubleDown(game)
		}
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(game)
		}
	}

}
