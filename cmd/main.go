package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"superheroesAPI/pkg/db"
	"superheroesAPI/pkg/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
	r.Get("/authenticate", h.Authenticate)
	r.Post("/logout", h.Logout)
	r.Get("/list", h.ListHeroes)
	r.Post("/create", h.CreateHero)
	r.Route("/{userId}/{id}", func(r chi.Router) {
		r.Get("/", h.GetHero)
		r.Post("/add", h.AddHeroToTeam)
		r.Delete("/delete", h.DeleteHero)
		r.Patch("/update", h.UpdateHero)
	})

	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Server is listening on port 4000...")

}
