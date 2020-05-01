package models

type BrokenDealer struct {
	Name      string
	Difficuty Difficuty
	Hand      *Hand
}

func (broken *BrokenDealer) Play(currentGame Game) Game {
	if broken.Hand.Score <= 17 &&
		currentGame.User.Hand.Score > currentGame.User.Hand.Score {
		return Hit(currentGame.Id, true)
	} else {
		return Stand(currentGame.Id)
	}
}
