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

func (h Handler) RemoveHeroFromTeam(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
	heroId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	var hero models.Hero

	tx := h.DB.Table("teamMembers").
		Where(fmt.Sprintf("user_id = %d AND hero_id = %d", userId, heroId)).
		Delete(hero)
	if tx.Error != nil {
		log.Println(tx.Error)
		return
	}

	err = json.NewEncoder(w).Encode("Hero removed from team")
	if err != nil {
		log.Println(err)
		return
	}
}
