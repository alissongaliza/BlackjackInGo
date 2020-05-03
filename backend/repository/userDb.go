package models

import (
	"fmt"
	"sync"

	"github.com/alissongaliza/BlackjackInGo/utils"
)

type UserDb map[int]User

var userInstance UserDb

var userOnce sync.Once

func GetUserDb() UserDb {
	userOnce.Do(func() {

		userInstance = make(UserDb)

		hand1 := NewHand()
		hand2 := NewHand()
		userInstance[1] = User{Player{&hand1}, "alisson", 1, 18, 100}
		userInstance[2] = User{Player{&hand2}, "a", 2, 22, 100}

	})

	return userInstance
}

func assignUserId(user *User) {
	users := GetUserDb()
	if users == nil {
		user.Id = 0
	}

	keys := utils.GetMapIntKeys(users)
	user.Id = utils.FindMaxIndex(keys)
}

func (users UserDb) Create(newUser User) User {
	assignUserId(&newUser)
	users[newUser.Id] = newUser
	return newUser
}

func IsUserValid(userId int) bool {
	users := GetUserDb()

	if user, present := users[userId]; present && user.Age >= 18 {
		return true
	}
	return false
}

func (db UserDb) Get(id int) User {
	users := GetUserDb()

	user, ok := users[id]
	if !ok {
		panic(fmt.Sprintf("User of id %d not found", id))
	}
	return user
}

func (db UserDb) List(username string) []User {
	users := GetUserDb()
	filteredUsers := make([]User, 0)
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

func (db *UserDb) Update(user User) User {
	users := GetUserDb()

	users[user.Id] = user
	return user
}
