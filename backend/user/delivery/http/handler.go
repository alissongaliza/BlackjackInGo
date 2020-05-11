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

type UserHandler struct {
	userUsecase user.UseCase
	gameUsecase game.UseCase
}

func NewUserHandler(r chi.Router, userUsecase user.UseCase, gameUsecase game.UseCase) {
	handler := &UserHandler{userUsecase: userUsecase, gameUsecase: gameUsecase}
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.listUsers)
		r.Post("/", handler.addUser)
		r.Route("/{id}", func(r chi.Router) {
			// 	r.Get("/", handler.findUser)
			// 	r.Post("/play", handler.hit)
			r.Get("/games", handler.listUserGames)
		})
	})
}

func (uh *UserHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listing users...")
	usernames, ok := r.URL.Query()["username"]
	if !ok {
		usernames[0] = ""
	}
	name := usernames[0]
	users := uh.userUsecase.ListUser(name)
	json.NewEncoder(w).Encode(users)
}

func (uh *UserHandler) listUserGames(w http.ResponseWriter, r *http.Request) {
	userId := utils.StringToInt(chi.URLParam(r, "id"))
	games := uh.gameUsecase.ListGame(userId)
	json.NewEncoder(w).Encode(games)
}

// func (uh *UserHandler) findUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "finding users...")
// }

func (uh *UserHandler) addUser(w http.ResponseWriter, r *http.Request) {
	var data utils.UserCreateRequest
	json.NewDecoder(r.Body).Decode(&data)

	newUser := uh.userUsecase.CreateUser(data.Name, data.Age)
	json.NewEncoder(w).Encode(newUser)

}
