package repositories

import (
	"fmt"
	"sync"

	"github.com/alissongaliza/BlackjackInGo/backend/game"
	"github.com/alissongaliza/BlackjackInGo/backend/models"
	"github.com/alissongaliza/BlackjackInGo/utils"
)

type inMemoryGameRepository struct {
	Db map[int]models.Game
}

var gameInstance inMemoryGameRepository

var once sync.Once

func NewInMemoryGameDb() game.Repository {
	once.Do(func() {

		gameInstance.Db = make(map[int]models.Game)

		// userDb := GetUserDb()
		// user := userDb.Get(2)
		// dealerHand := NewHand()

		// dealer := models.Dealer{&EasyDealer{}, models.Player{&dealerHand}, utils.Easy}
		// cards := NewDeck()
		// gameInstance.Db[1] = models.Game{1, user, dealer, cards, 0, 0, utils.NoAction, utils.NoAction, utils.Setup}

	})

	return &gameInstance
}

func (imgr *inMemoryGameRepository) GetNextValidGameId() int {
	games := imgr.Db
	if games == nil {
		return 1
	}

	keys := utils.GetMapIntKeys(games)
	return utils.FindMaxIndex(keys) + 1
}

func (imgr *inMemoryGameRepository) CreateGame(newGame models.Game) models.Game {
	games := imgr.Db
	games[newGame.Id] = newGame
	return newGame
}

func (imgr *inMemoryGameRepository) IsGameValid(gameId int) bool {
	games := imgr.Db
	return len(games) > 0
}

func (imgr *inMemoryGameRepository) GetGame(id int) models.Game {
	game, ok := imgr.Db[id]
	if !ok {
		panic(fmt.Sprintf("Game of id %d not found", id))
	}
	return game
}

func (imgr *inMemoryGameRepository) UpdateGame(game models.Game) models.Game {
	games := imgr.Db
	games[game.Id] = game
	return game
}

func (imgr *inMemoryGameRepository) ListGame(userId int) []models.Game {
	games := imgr.Db
	filteredGames := make([]models.Game, 0)
	//this whole function is embarassing
	if userId == 0 {
		for _, game := range games {
			filteredGames = append(filteredGames, game)
		}
	} else {
		for _, game := range games {
			if game.User.Id == userId &&
				game.GameState == utils.Playing ||
				game.GameState == utils.Setup {
				filteredGames = append(filteredGames, game)
			}
		}
	}
	return filteredGames
}

func (imgr *inMemoryGameRepository) CreateHand() models.Hand {
	return models.Hand{Cards: make([]models.Card, 0), Score: 0}
}

func (imgr *inMemoryGameRepository) CreateDeck() (newDeck []models.Card) {
	for _, suit := range []utils.SuitType{utils.Hearts, utils.Spades, utils.Clubs, utils.Diamonds} {
		newDeck = append(newDeck,
			models.Card{Suit: suit, Name: "2", IsNumber: true, IsFaceUp: false},
			models.Card{Suit: suit, Name: "3", IsNumber: true, IsFaceUp: false},
			models.Card{Suit: suit, Name: "4", IsNumber: true, IsFaceUp: false},
			models.Card{Suit: suit, Name: "5", IsNumber: true, IsFaceUp: false},
			models.Card{Suit: suit, Name: "6", IsNumber: true, IsFaceUp: false},
			models.Card{Suit: suit, Name: "8", IsNumber: true, IsFaceUp: false},
			models.Card{Suit: suit, Name: "7", IsNumber: true, IsFaceUp: false},
			models.Card{Suit: suit, Name: "9", IsNumber: true, IsFaceUp: false},
			models.Card{Suit: suit, Name: "10", IsNumber: true, IsFaceUp: false},
			models.Card{Suit: suit, Name: "J", IsNumber: false, IsFaceUp: false},
			models.Card{Suit: suit, Name: "Q", IsNumber: false, IsFaceUp: false},
			models.Card{Suit: suit, Name: "K", IsNumber: false, IsFaceUp: false},
			models.Card{Suit: suit, Name: "ACE", IsNumber: false, IsFaceUp: false},
		)
	}
	return
}
