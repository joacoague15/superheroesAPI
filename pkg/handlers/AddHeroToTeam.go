package handlers

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"superheroesAPI/pkg/models"
)

func (h Handler) AddHeroToTeam(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
	heroId, err := strconv.ParseInt(chi.URLParam(r, "heroId"), 10, 64)

	if err != nil {
		log.Println(err)
	}

	teamMember := models.TeamMember{
		UserId: userId,
		HeroId: heroId,
	}

	tx := h.DB.Table("teamMembers").Create(&teamMember)
	if tx.Error != nil {
		log.Println(tx.Error)
	}
}
