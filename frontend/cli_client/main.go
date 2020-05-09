package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alissongaliza/BlackjackInGo/cliClient/api"
	gamePkg "github.com/alissongaliza/BlackjackInGo/cliClient/game"
	"github.com/alissongaliza/BlackjackInGo/cliClient/utils"
	remoteUtils "github.com/alissongaliza/BlackjackInGo/utils"
)

func main() {
	fmt.Print("Welcome to the ultimate BlackJack made entirely in Golang!\n")
	fmt.Println("Do you have an account already?\n1- Yes \n2- No")
	var answer int
	fmt.Scan(&answer)
	var user utils.User
	if answer == 1 {
		fmt.Println("Great! What's your username?")
		username, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println("Alright, hold on a second...")

		// users := api.FindUser("a")
		// trimming new line
		username = username[:len(username)-1]
		user = api.FindUser(username)
	} else {
		// CreateUser(user)
		fmt.Println("Create a account. Let's start with you username: ")
		newUsername, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		newUsername = newUsername[:len(newUsername)-1]
		fmt.Println("How old are you?")
		var age int
		fmt.Scan(&age)
		user = api.CreateUser(age, newUsername)
	}

	fmt.Printf("We're all set, %s! Want to start a new game "+
		"or continue a previous one? \n1- New Game \n2- Continue where I left off\n", user.Name)
	fmt.Scan(&answer)
	var game utils.Game
	if answer == 1 {
		//create game
		diff := chooseDifficulty()
		betAnswer := placeBets(user)
		game = api.CreateGame(user.Id, diff, betAnswer)
	} else {
		// find ongoing games
		games := api.FindOngoingGamesOfUser(user.Id)
		fmt.Println("Choose which game you want to continue on:")
		game = printOngoingGames(games)
	}
	fmt.Println("\nStarting game...")
	gamePkg.EnterGameLoop(game)
}

func placeBets(user utils.User) int {
	var betAnswer int
	fmt.Println("How much do you like to bet? Minimum of 25")
	fmt.Println("Current balance: ", user.Chips)
	for {
		fmt.Scan(&betAnswer)
		if betAnswer > user.Chips {
			fmt.Println("Can't bet more than your current balance. Try again.")
		} else if betAnswer < 25 {
			fmt.Println("Minimum bet is 25. Try again.")
		} else {
			return betAnswer
		}
	}
}

func chooseDifficulty() remoteUtils.Difficulty {
	var diffAnswer int
	fmt.Println("What difficulty you want to play on?\n1- Easy\n2- Broken")
	for {
		fmt.Scan(&diffAnswer)
		if diffAnswer == 1 {
			return remoteUtils.Easy
		} else if diffAnswer == 2 {
			return remoteUtils.Broken
		} else {
			fmt.Println("Invalid Option. Try again")
		}
	}
}

func printOngoingGames(games []utils.Game) utils.Game {
	for i, game := range games {
		fmt.Printf("%d- ", i+1)
		gamePkg.PrintTable(game)
	}
	var pickedGame int
	fmt.Scan(&pickedGame)
	// started with 1
	return games[pickedGame-1]
}
