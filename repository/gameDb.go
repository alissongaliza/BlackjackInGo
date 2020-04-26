package models

import (
	"fmt"
	"sync"

	"github.com/alissongaliza/BlackjackInGo/utils"
)

type GameDb map[int]Game

var gameInstance GameDb

var gameOnce sync.Once

func GetGameDb() GameDb {
	gameOnce.Do(func() {

		gameInstance = make(GameDb)

		userDb := GetUserDb()
		user := userDb.Get(1)
		dealerHand := NewHand()
		opponentName := fmt.Sprintf("%s's opponent", user.Name)
		dealer := EasyDealer{opponentName, easy, &dealerHand}
		cards := NewDeck()
		gameInstance[1] = Game{1, user, &dealer, cards, 0, false, noAction, noAction, playing}

	})

	return gameInstance
}

func assignGameId(game *Game) {
	games := GetGameDb()
	if games == nil {
		game.Id = 0
	}

	keys := utils.GetMapIntKeys(games)
	game.Id = utils.FindMaxIndex(keys)
}

func (games GameDb) Create(newGame Game) Game {
	assignGameId(&newGame)
	games[newGame.Id] = newGame
	return newGame
}

func IsGameValid(gameId int) bool {
	games := GetGameDb()

	if _, present := games[gameId]; present {
		return true
	}
	return false
}

func (games GameDb) Get(id int) Game {
	game, ok := games[id]
	if !ok {
		panic(fmt.Sprintf("Game of id %d not found", id))
	}
	return game
}

func (games GameDb) Update(game Game) Game {
	games[game.Id] = game
	return game
}

func getDealer(dif Difficuty, opponentName string) (dealer Dealer) {
	hand := NewHand()
	switch dif {
	case easy:
		dealer = &EasyDealer{opponentName, easy, &hand}
	case medium:
		dealer = &MediumDealer{opponentName, medium, &hand}
	case hard:
		dealer = &HardDealer{opponentName, hard, &hand}
	}
	return
}
