package models

type User struct {
	Name string
	Id   int
	Age  int
}

const UserConst Actors = "user"

func NewUser(name string, age int) (newUser User) {

	newUser = User{name, -1, age}
	addNewUser(&newUser)
	return newUser
}

// func (user *User) SetId(id int) {
// 	user.Id = id
// }

// func (user *User) SetName(name string) {
// 	user.Name = name
// }

// func (user User) GetId() int {
// 	return user.Id
// }

// func (user User) GetName() string {
// 	return user.Name
// }
