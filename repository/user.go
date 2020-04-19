package models

type User interface {
	hit(game int)
	stand(game int)
	doubleDown(game int)
}
