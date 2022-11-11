package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"superheroesAPI/pkg/models"
)

func (h Handler) AddHeroToTeam(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
	heroId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		_ = json.NewEncoder(w).Encode("Error adding team member")
		w.WriteHeader(http.StatusInternalServerError)
	}

	teamMember := models.TeamMember{
		UserId: userId,
		HeroId: heroId,
	}

	err = json.NewEncoder(w).Encode("Team hero added")

	tx := h.DB.Table("teamMembers").Debug().Create(&teamMember)
	if tx.Error != nil {
		_ = json.NewEncoder(w).Encode("Error adding team member")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
