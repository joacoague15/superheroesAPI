package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"superheroesAPI/pkg/models"
)

func (h Handler) DeleteHero(w http.ResponseWriter, r *http.Request) {
	var hero models.Hero
	id := chi.URLParam(r, "id")

	tx := h.DB.Table("heroes").
		Where(fmt.Sprintf("heroes.id = %s", id)).
		Delete(hero)
	if tx.Error != nil {
		log.Println(tx.Error)
		return
	}

	err := json.NewEncoder(w).Encode("Hero deleted")
	if err != nil {
		log.Println(err)
		return
	}
}
