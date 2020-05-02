package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alissongaliza/BlackjackInGo/utils"

	models "github.com/alissongaliza/BlackjackInGo/backend/repository"

	"github.com/go-chi/chi"
)

type userRequest struct {
	*models.User
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listing users...")
	username, ok := r.URL.Query()["username"]
	var users []models.User
	if !ok {
		username[0] = ""
	}
	users = models.GetUserDb().List(username[0])
	fmt.Println(users[0])
	json.NewEncoder(w).Encode(users)

}

func findUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "finding users...")
}

func addUser(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			// fmt.Fprint(w, r)
			fmt.Println(r)
		}
	}()
	var data models.User
	json.NewDecoder(r.Body).Decode(&data)

	newUser := models.NewUser(data.Name, data.Age)
	newUser = models.GetUserDb().Create(newUser)
	fmt.Println(newUser)
	json.NewEncoder(w).Encode(newUser)
	// fmt.Fprintf(w, "adding user...")
	// fmt.Fprintf(w, "%+v", newUser)

}

func hit(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			// fmt.Fprint(w, r)
			fmt.Println(r)
		}
	}()
	fmt.Println("hit called")
	playerId := utils.StringToInt(chi.URLParam(r, "id"))
	var hitRequest struct {
		GameId int
	}
	fmt.Println("hit", hitRequest)
	json.NewDecoder(r.Body).Decode(&hitRequest)
	player := models.GetUserDb().Get(playerId)
	game := player.Hit(hitRequest.GameId, true)
	json.NewEncoder(w).Encode(game)

}

func stand(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			// fmt.Fprint(w, r)
			fmt.Println(r)
		}
	}()
	fmt.Println("stand called")
	playerId := utils.StringToInt(chi.URLParam(r, "id"))
	var gameId int
	json.NewDecoder(r.Body).Decode(&gameId)
	user := models.GetUserDb().Get(playerId)
	game := user.Stand(gameId)
	json.NewEncoder(w).Encode(game)

}

func doubleDown(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			// fmt.Fprint(w, r)
			fmt.Println(r)
		}
	}()
	fmt.Println("doubleDown called")
	playerId := utils.StringToInt(chi.URLParam(r, "id"))
	var gameId int
	json.NewDecoder(r.Body).Decode(&gameId)
	player := models.GetUserDb().Get(playerId)
	game := player.DoubleDown(gameId)
	json.NewEncoder(w).Encode(game)

}

func UserRouter() chi.Router {
	// func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", listUsers)
	r.Post("/", addUser)
	r.Get("/{id}", findUser)
	r.Post("/{id}/hit", hit)
	r.Post("/{id}/stand", stand)
	r.Post("/{id}/doubleDown", doubleDown)

	return r
}
