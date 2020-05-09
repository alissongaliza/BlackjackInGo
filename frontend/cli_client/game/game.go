package game

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/cliClient/api"
	"github.com/alissongaliza/BlackjackInGo/cliClient/utils"
	remoteUtils "github.com/alissongaliza/BlackjackInGo/utils"
)

func EnterGameLoop(game utils.Game) {
	if game.GameState == remoteUtils.Setup {
		game = api.InitGame(game.Id)
	}
	for game.GameState == remoteUtils.Playing {
		var answer int
		PrintTable(game)
		printAvailableOptions(game)
		fmt.Scan(&answer)
		switch answer {
		//for simplicity options numbers are static
		case 1:
			{
				api.PlayAction(remoteUtils.Hit, &game)
			}
		case 2:
			{
				api.PlayAction(remoteUtils.Stand, &game)
			}
		case 3:
			{
				api.PlayAction(remoteUtils.DoubleDown, &game)
			}
		default:
			{
				fmt.Println("Invalid option.")
			}
		}
	}
	PrintTable(game)
	fmt.Printf("Game is over, you %s\n", game.GameState)
	printPayout(game)
	fmt.Println("\nCurrent balance is ", game.User.Chips)
}

func printAvailableOptions(game utils.Game) {
	user := game.User
	fmt.Println("Choose your action")
	//hit
	if user.Hand.Score < 21 {
		fmt.Println("1- Hit")
	}
	//stand (he can always stand)
	fmt.Println("2- Stand")
	//doubleDown
	if len(user.Hand.Cards) == 2 && user.Hand.Score < 21 {
		fmt.Println("3- Double Down")
	}

}

func PrintTable(game utils.Game) {
	fmt.Print("Your hand: ")
	userHand := game.User.Hand
	PrintHand(*userHand)
	fmt.Print("Dealer hand: ")
	dealerHand := game.Dealer.Hand
	PrintHand(*dealerHand)
}

func PrintHand(hand utils.Hand) {
	if len(hand.Cards) == 0 {
		fmt.Println("Empty")
		return
	}
	userCards := ""
	for _, card := range hand.Cards {
		if card.IsFaceUp {
			userCards += " " + card.Name + " " + string(card.Suit) + ","
		} else {
			// card faced down
			userCards += " ?,"
		}
	}
	// switch comma for period
	userCards = userCards[:len(userCards)-1]
	userCards += fmt.Sprintf(". Score: %d.", hand.Score)
	fmt.Println(userCards)
}

func printPayout(game utils.Game) {
	if game.GameState == remoteUtils.Lost {
		fmt.Printf("You lost the chips you bet (%d)", game.Bet)
	} else if game.GameState == remoteUtils.Won {
		fmt.Printf("You won %d chips", game.Payout)
	} else {
		fmt.Printf("You drew and got %d chips", game.Payout)
	}
}
