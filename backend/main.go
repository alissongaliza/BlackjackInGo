package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	dealerUsecases "github.com/alissongaliza/BlackjackInGo/backend/dealer/usecases"
	gameDeliveryHttp "github.com/alissongaliza/BlackjackInGo/backend/game/delivery/http"
	gameRepositories "github.com/alissongaliza/BlackjackInGo/backend/game/repositories"
	gameUsecases "github.com/alissongaliza/BlackjackInGo/backend/game/usecases"
	userDeliveryHttp "github.com/alissongaliza/BlackjackInGo/backend/user/delivery/http"
	userRepositories "github.com/alissongaliza/BlackjackInGo/backend/user/repositories"
	userUsecases "github.com/alissongaliza/BlackjackInGo/backend/user/usecases"
)

func pingLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", pingLink)
	userRepo := userRepositories.NewInMemoryUserDb()
	gameRepo := gameRepositories.NewInMemoryGameDb()
	dealerUsescase := dealerUsecases.NewDealerUsecase(gameRepo, userRepo)
	userUsecases := userUsecases.NewUserUsecase(gameRepo, userRepo, dealerUsescase)
	gameUsecases := gameUsecases.NewGameUsecase(gameRepo, dealerUsescase, userUsecases, userRepo)
	userDeliveryHttp.NewUserHandler(r, userUsecases, gameRepo, userRepo)
	gameDeliveryHttp.NewGameHandler(r, gameUsecases, userUsecases, userRepo, gameRepo)

	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(":8080", r))
}
