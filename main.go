package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alissongaliza/BlackjackInGo/game"
	"github.com/alissongaliza/BlackjackInGo/player"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func pingLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

// func initInMemoryDataStore() map[Actors]interface{} {

// 	// defer models.Sess.Close()
// 	// models.Sess.Exec(models.)
// 	db = models.GetDb()
// 	return db
// }

func main() {
	// db := initInMemoryDataStore()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", pingLink)
	r.Mount("/players", player.PlayerRouter())
	r.Mount("/games", game.GameRouter())
	// r.Mount(game.gameRoutes)

	fmt.Println("Listening")

	// r.HandleFunc("/players", pingLink)
	log.Fatal(http.ListenAndServe(":8080", r))
}
