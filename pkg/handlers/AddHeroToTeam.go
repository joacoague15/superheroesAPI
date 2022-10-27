package handlers

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"superheroesAPI/pkg/models"
)

func (h Handler) AddHeroToTeam(_ http.ResponseWriter, r *http.Request) {
	heroId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		log.Println(err)
	}

	teamMember := models.TeamMember{
		// HARDCODED
		UserId: 8,
		HeroId: heroId,
	}

	tx := h.DB.Table("teamMembers").Debug().Create(&teamMember)
	if tx.Error != nil {
		log.Println(tx.Error)
	}
}
