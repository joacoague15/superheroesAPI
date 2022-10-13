package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"superheroesAPI/pkg/models"
)

func (h handler) GetHeroes(w http.ResponseWriter, _ *http.Request) {
	var heroes []models.Hero

	tx := h.DB.Table("heroes").Select("*").Limit(100)
	tx.Find(&heroes)
	if tx.Error != nil {
		log.Println(tx.Error)
	}

	err := json.NewEncoder(w).Encode(heroes)
	if err != nil {
		log.Println(err)
	}
}
