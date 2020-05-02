package main

import (
	"fmt"

	"github.com/alissongaliza/BlackjackInGo/cliClient/api"
	gamePkg "github.com/alissongaliza/BlackjackInGo/cliClient/game"
	"github.com/alissongaliza/BlackjackInGo/cliClient/utils"
	remoteUtils "github.com/alissongaliza/BlackjackInGo/utils"
)

func main() {
	fmt.Print("Welcome to the ultimate BlackJack made entirely in Golang!\n")
	fmt.Println("Do you have an account already?\n1- Yes \n2- No")
	// var answer int
	// fmt.Scan(&answer)
	var user utils.User
	// if answer == 1 {
	fmt.Println("Great! What's your username?")
	// username, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Alright, hold on a second...")

	users := api.FindUser("a")
	// users := api.FindUser(username[:len(username)-1])
	user = users[0]
	// CreateUser(user)
	// } else {

	// }

	fmt.Printf("We're all set, %s! Want to start a new game "+
		"or continue a previous one? \n1- New Game \n2- Continue where I left off\n", user.Name)
	// fmt.Scan(&answer)
	var game utils.Game
	// if answer == 1 {
	//create game
	// fmt.Println("What difficulty you want to play on?\n 1- Easy\n2- Broken")
	// fmt.Scan(&answer)
	game = api.StartGame(user.Id, remoteUtils.Easy, 25)
	// } else {
	//find all ongoing games related to this user
	// }
	fmt.Println("game:", game.Id)
	gamePkg.EnterGameLoop(game)
}
