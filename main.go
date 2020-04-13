package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alissongaliza/BlackjackInGo/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func pathPrinter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func pingLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", pingLink)
	r.Mount("/users", user.UserRouter())
	// r.Mount(game.gameRoutes)

	fmt.Println("Listening")

	// r.HandleFunc("/users", pingLink)
	log.Fatal(http.ListenAndServe(":8080", r))
}
