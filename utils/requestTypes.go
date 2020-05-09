package utils

type UserCreateRequest struct {
	Name string
	Age  int
}

type UserActionRequest struct {
	Action Action
}

type GameCreateRequest struct {
	UserId int
	Dif    Difficulty
	Bet    int
}
