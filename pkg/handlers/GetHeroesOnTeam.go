package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"superheroesAPI/pkg/models"
)

func (h Handler) GetHeroesOnTeam(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)

	var teamHeroes []models.Hero

	tx := h.DB.Table("teamMembers").Debug().
		Select("*").Limit(3).
		Joins(`LEFT JOIN heroes ON heroes.id = "teamMembers".hero_id`).
		Where(fmt.Sprintf(`"teamMembers".user_id = %d`, userId))
	tx.Find(&teamHeroes)

	if tx.Error != nil {
		log.Println(tx.Error)
	}

	err = json.NewEncoder(w).Encode(teamHeroes)
	if err != nil {
		log.Println(err)
	}
}
