package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"superheroesAPI/pkg/models"
)

func (h Handler) UpdateHero(w http.ResponseWriter, r *http.Request) {
	var hero models.Hero
	id := chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&hero)

	hero.Id, err = strconv.ParseInt(id, 10, 64)

	tx := h.DB.Table("heroes").Save(&hero)
	if tx.Error != nil {
		log.Println(tx.Error)
		return
	}

	err = json.NewEncoder(w).Encode("Hero updated")
	if err != nil {
		log.Println(err)
		return
	}
}
