# BlackjackInGo

## The sole purpose of this project is to learn something new, in this case, Golang

### The focus of this project

-   Get familiar with Golang syntax
-   Learn the thought process of the language
-   Understand the idiomatic ways of the language
-   Integrate the project with third-party libraries

### BlackjackInGo

This project implements a simplistic Blackjack game.

#### How does it work

- 1 deck (4 suits, 52 cards)
- Actions: Hit, Stand and Double down
- Dealer: Easy and Broken difficulties
- Singleplayer
- You can start another game at any time and resume previous games later

#### Stack

- [Docker (compose)](https://docs.docker.com/compose/compose-file/)
- [Go](https://golang.org)
- [go-chi lib](https://github.com/go-chi/chi)
- [CompileDaemon](https://github.com/githubnemo/CompileDaemon)

#### How to run

1.  Initializing server container (using port 8080)
    ```bash
    docker-compose up
    ```
2.  Run cli client
    ```bash
    go run frontend/cli_client/main.go
    ```
3.  Have fun 

### Resources used so far

- [The Zoo of Go Functions](https://blog.learngoprogramming.com/go-functions-overview-anonymous-closures-higher-order-deferred-concurrent-6799008dde7b)
- [Go standart lib](https://golang.org/doc/)
- [Go tour](https://tour.golang.org/methods/)
- [Everything you need to know about packages in Go](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)
- [Interfaces in Go](https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c)
- [Pointers in Go](https://www.callicoder.com/golang-pointers/)
- [Clean Architecture repo example](https://github.com/bxcodec/go-clean-arch/tree/v2)
- [Golang Clean Architecture](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)
- [How singleton pattern works with Golang](https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f)
- [Unit testing made easy in go](https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318)
- [Blackjack rules](https://www.pagat.com/banking/blackjack.html)
- and **lots** of [stackoverflow](https://stackoverflow.com/questions/tagged/go) discussions
