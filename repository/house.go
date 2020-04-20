package models

import "fmt"

type Difficuty string

const (
	easy   Difficuty = "easy"
	medium Difficuty = "medium"
	hard   Difficuty = "hard"
)

type House interface {
	User
}

type EasyHouse struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

type MediumHouse struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

type HardHouse struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

func (easy EasyHouse) hit(gameId int, faceUp bool) {
	fmt.Println("player hit!")
}

func (easy EasyHouse) stand(gameId int) {
	fmt.Println("player stand!")
}

func (easy EasyHouse) doubleDown(gameId int) {
	fmt.Println("player doubleDown!")
}

func (medium MediumHouse) hit(gameId int, faceUp bool) {
	fmt.Println("player hit!")
}
func (medium MediumHouse) stand(gameId int) {
	fmt.Println("player stand!")
}
func (medium MediumHouse) doubleDown(gameId int) {
	fmt.Println("player doubleDown!")
}

func (hard HardHouse) hit(gameId int, faceUp bool) {
	fmt.Println("player hit!")
}

func (hard HardHouse) stand(gameId int) {
	fmt.Println("player stand!")
}

func (hard HardHouse) doubleDown(gameId int) {
	fmt.Println("player doubleDown!")
}
