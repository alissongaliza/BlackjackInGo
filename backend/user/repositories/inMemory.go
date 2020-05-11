package repositories

import (
	"fmt"
	"sync"

	"github.com/alissongaliza/BlackjackInGo/backend/user"

	"github.com/alissongaliza/BlackjackInGo/backend/models"
	"github.com/alissongaliza/BlackjackInGo/utils"
)

type inMemoryUserRepository struct {
	Db map[int]models.User
}

var userInstance inMemoryUserRepository

var once sync.Once

func NewInMemoryUserDb() user.Repository {
	once.Do(func() {

		userInstance.Db = make(map[int]models.User)

		// hand1 := game.UseCase.GetNewHand()
		// hand2 := game.UseCase.GetNewHand()
		// userInstance.Db[1] = models.User{models.Player{hand1}, "alisson", 1, 18, 100}
		// userInstance.Db[2] = models.User{models.Player{hand2}, "a", 2, 22, 100}

	})

	return &userInstance
}

func (imur *inMemoryUserRepository) GetNextValidUserId() int {
	users := imur.Db
	if users == nil {
		return 1
	}

	keys := utils.GetMapIntKeys(users)
	return utils.FindMaxIndex(keys) + 1
}

func (imur *inMemoryUserRepository) CreateUser(user models.User) models.User {
	users := imur.Db
	users[user.Id] = user
	return user
}

func (imur *inMemoryUserRepository) GetUser(id int) models.User {
	users := imur.Db

	user, ok := users[id]
	if !ok {
		panic(fmt.Sprintf("User of id %d not found", id))
	}
	return user
}

func (imur *inMemoryUserRepository) ListUser(username string) []models.User {
	users := imur.Db
	filteredUsers := make([]models.User, 0)
	//this whole function is embarassing
	if username == "" {
		for _, user := range users {
			filteredUsers = append(filteredUsers, user)
		}
	} else {
		for _, user := range users {
			if user.Name == username {
				filteredUsers = append(filteredUsers, user)
				break
			}
		}
	}
	return filteredUsers
}

func (imur *inMemoryUserRepository) UpdateUser(user models.User) models.User {
	users := imur.Db

	users[user.Id] = user
	return user
}
