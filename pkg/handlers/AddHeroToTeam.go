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

func (h Handler) AddHeroToTeam(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
	heroId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	var teamHeroes []models.Hero
	var heroesCount int64

	if err != nil {
		_ = json.NewEncoder(w).Encode("Error adding team member")
		w.WriteHeader(http.StatusInternalServerError)
	}

	tx := h.DB.Table("teamMembers").Debug().
		Select("*").Limit(3).
		Joins(`LEFT JOIN heroes ON heroes.id = "teamMembers".hero_id`).
		Where(fmt.Sprintf(`"teamMembers".user_id = %d`, userId))
	tx.Find(&teamHeroes).
		Count(&heroesCount)

	heroIsAlreadyOnTeam(heroId, teamHeroes)

	if tx.Error != nil {
		log.Println(tx.Error)
	}

	teamMember := models.TeamMember{
		UserId: userId,
		HeroId: heroId,
	}

	if heroesCount > 2 {
		err = json.NewEncoder(w).Encode("You cannot have more than 3 heroes on your team")
	} else if heroIsAlreadyOnTeam(heroId, teamHeroes) {
		err = json.NewEncoder(w).Encode("You cannot have same hero twice")
	} else {
		err = json.NewEncoder(w).Encode("Team hero added")

		tx = h.DB.Table("teamMembers").Debug().Create(&teamMember)
		if tx.Error != nil {
			_ = json.NewEncoder(w).Encode("Error adding team member")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

}

func heroIsAlreadyOnTeam(idHero int64, heroesOnTeam []models.Hero) bool {
	var isOnTheTeam = false

	for _, hero := range heroesOnTeam {
		if hero.Id == idHero {
			isOnTheTeam = true
		}
	}

	return isOnTheTeam
}
