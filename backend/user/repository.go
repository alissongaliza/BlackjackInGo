package user

import "github.com/alissongaliza/BlackjackInGo/backend/models"

type Repository interface {
	CreateUser(user models.User) models.User
	GetUser(userId int) models.User
	ListUser(name string) []models.User
	UpdateUser(user models.User) models.User
	GetNextValidUserId() int
}
