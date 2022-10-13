package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"superheroesAPI/pkg/db"
	"superheroesAPI/pkg/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", h.GetHeroes)
	r.Post("/create", h.CreateHero)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.GetHero)
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
