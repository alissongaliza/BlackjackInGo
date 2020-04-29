package main

import (
	"fmt"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	user := PostUser{"asdasd", 18}
	CreateUser(user)
	// text, _ := reader.ReadString('\n')
	// fmt.Println("text", text)

}
