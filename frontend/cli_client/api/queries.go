package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/alissongaliza/BlackjackInGo/cliClient/utils"
	remoteUtils "github.com/alissongaliza/BlackjackInGo/utils"
)

func CreateUser(age int, name string) (newUser utils.User) {

	userRequest := utils.NewUserRequest{Age: age, Name: name}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&userRequest)
	req, err := http.NewRequest("POST", "http://localhost:8080/users", buf)
	if err != nil {
		log.Print(err)
	}

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Print(e)
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&newUser)
	fmt.Println(newUser)
	return
}

func FindUser(username string) utils.User {

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
	var users []utils.User
	json.NewDecoder(res.Body).Decode(&users)

	if len(users) > 0 {
		return users[0]
	} else {
		panic("No user found with this username")
	}
}

func FindOngoingGamesOfUser(userId int) []utils.Game {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&userId)
	req, err := http.NewRequest("GET", "http://localhost:8080/games", buf)
	q := req.URL.Query()
	q.Add("userId", strconv.Itoa(userId))
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
	var games []utils.Game
	json.NewDecoder(res.Body).Decode(&games)

	return games
}

func InitGame(gameId int) utils.Game {

	buf := new(bytes.Buffer)
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/games/%d", gameId), buf)
	if err != nil {
		log.Print(err)
	}

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Print(e)
	}
	defer res.Body.Close()
	var game utils.Game
	json.NewDecoder(res.Body).Decode(&game)
	return game
}

func CreateGame(userId int, dif remoteUtils.Difficulty, bet int) (game utils.Game) {

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
	json.NewEncoder(buf).Encode(remoteUtils.UserActionRequest{Action: action})
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8080/games/%d/play", game.Id), buf)
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

}
