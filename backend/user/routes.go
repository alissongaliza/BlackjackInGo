package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alissongaliza/BlackjackInGo/backend/utils"

	models "github.com/alissongaliza/BlackjackInGo/backend/repository"

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
	json.NewDecoder(r.Body).Decode(&hitRequest)
	player := models.GetUserDb().Get(playerId)
	player.Hit(hitRequest.GameId, true)
	// json.NewEncoder(w).Encode(newUser)

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
	user.Stand(gameId)

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
	player.DoubleDown(gameId)

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
