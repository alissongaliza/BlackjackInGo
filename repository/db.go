package models

import (
	"sync"
)

// var settings = postgresql.ConnectionURL{
// 	Host:     "db",
// 	Database: "BlackJackDB",
// 	User:     "bj",
// 	Password: "bjIsFun",
// }

// var Sess, Err = postgresql.Open(settings)
type DB map[Actors]map[int]interface{}

type Actors string

var once sync.Once

var instance DB

func GetDb() DB {
	once.Do(func() {

		instance = make(DB, 2)
		instance[UserConst] = make(map[int]interface{})
		instance[GameConst] = make(map[int]interface{})
		addNewUser(&User{"Alisson", -1, 18})

	})

	return instance
}

func getMapKeys(m map[int]interface{}) []int {

	keys := make([]int, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func findNextId(key Actors) int {
	db := GetDb()
	if db == nil {
		return -1
	}

	collection := db[key]
	keys := getMapKeys(collection)

	// find max index
	max := 0
	for k, e := range keys {
		if k == 0 || e > max {
			max = e
		}
	}
	max += 1
	return max
}

func addNewUser(newUser *User) bool {
	db := GetDb()

	nextId := findNextId(UserConst)
	users := db[UserConst]
	newUser.Id = nextId
	users[nextId] = *newUser
	return true
}

func IsUserValid(userId int) bool {
	db := GetDb()

	if value, present := db[UserConst][userId].(User); present && value.Age >= 18 {
		return true
	}
	return false
}
