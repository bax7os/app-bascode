package controllers

import (
	"app/src/cookies"
	"net/http"
)

func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
