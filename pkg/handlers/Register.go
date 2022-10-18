package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"superheroesAPI/pkg/models"
)

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
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

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		return
	}

	tx := h.DB.Table("users").Create(&user)
	if tx.Error != nil {
		log.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("User created")
	if err != nil {
		log.Println(err)
		return
	}
}
