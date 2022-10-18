package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"superheroesAPI/pkg/models"
	"time"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
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

	var userSearched models.User
	var userFound models.User

	err = json.Unmarshal(body, &userSearched)
	if err != nil {
		log.Println(err)
		return
	}

	tx := h.DB.Table("users").
		Where("email = ?", userSearched.Email).
		First(&userFound)
	tx.Find(&userFound)

	if tx.Error != nil {
		log.Println(tx.Error)
		return
	}

	if userFound.Email == "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		err = json.NewEncoder(w).Encode("User not found")
		return
	}

	if userFound.Password != userSearched.Password {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		err = json.NewEncoder(w).Encode("Password is incorrect")
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: strconv.FormatInt(userFound.Id, 10),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(err)
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		log.Println(err)
		return
	}
}
