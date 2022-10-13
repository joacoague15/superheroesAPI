package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"superheroesAPI/pkg/models"
)

func (h handler) CreateHero(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(r.Body)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var hero models.Hero
	err = json.Unmarshal(body, &hero)
	if err != nil {
		log.Println(err)
		return
	}

	result := h.DB.Table("heroes").Create(&hero)
	if result.Error != nil {
		log.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Hero created")
	if err != nil {
		log.Println(err)
		return
	}
}
