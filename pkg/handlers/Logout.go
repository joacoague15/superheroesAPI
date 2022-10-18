package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (h Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	err := json.NewEncoder(w).Encode("Logout success")

	if err != nil {
		log.Println(err)
	}
}
