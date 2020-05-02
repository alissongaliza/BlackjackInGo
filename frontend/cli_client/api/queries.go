package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alissongaliza/BlackjackInGo/cliClient/utils"
	remoteUtils "github.com/alissongaliza/BlackjackInGo/utils"
)

func CreateUser(user utils.User) {

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(user)
	req, err := http.NewRequest("POST", "http://localhost:8080/user", buf)
	if err != nil {
		log.Print(err)
	}

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Print(e)
	}

	defer res.Body.Close()
}

func FindUser(username string) (users []utils.User) {

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(username)
	req, err := http.NewRequest("GET", "http://localhost:8080/users", buf)
	q := req.URL.Query()
	q.Add("username", username)
	req.URL.RawQuery = q.Encode()
	if err != nil {
		log.Print(err)
	}

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Print(e)
	}

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&users)

	return users
}

func StartGame(userId int, dif remoteUtils.Difficulty, bet int) (game utils.Game) {

	type startType struct {
		UserId int
		Dif    remoteUtils.Difficulty
		Bet    int
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&startType{Bet: bet, Dif: dif, UserId: userId})
	req, err := http.NewRequest("POST", "http://localhost:8080/games", buf)
	if err != nil {
		log.Print(err)
	}

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Print(e)
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&game)
	return
}

func PlayAction(action remoteUtils.Action, game *utils.Game) {

	buf := new(bytes.Buffer)
	type playRequest struct {
		GameId int
	}
	json.NewEncoder(buf).Encode(playRequest{GameId: game.Id})
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/users/%d/hit", game.User.Id), buf)
	if err != nil {
		log.Print(err)
	}

	client := &http.Client{}
	res, e := client.Do(req)
	defer res.Body.Close()
	if e != nil {
		log.Print(e)
	}

	json.NewDecoder(res.Body).Decode(&game)

}
