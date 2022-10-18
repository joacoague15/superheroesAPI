package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"superheroesAPI/pkg/models"
)

func (h Handler) GetHero(w http.ResponseWriter, r *http.Request) {
	var hero models.Hero
	id := chi.URLParam(r, "id")

	tx := h.DB.Table("heroes").
		Select("*").
		Where(fmt.Sprintf("heroes.id = %s", id))
	tx.Find(&hero)
	if tx.Error != nil {
		log.Println(tx.Error)
	}

	err := json.NewEncoder(w).Encode(hero)
	if err != nil {
		log.Println(err)
	}
}
