package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"superheroesAPI/pkg/models"
)

const SecretKey = "secret"

func (h Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("jwt")

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println(err)
		return
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	h.DB.Table("users").Where("id = ?", claims.Issuer).First(&user)

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		log.Println(err)
		return
	}

}
