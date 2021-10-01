package middlew

import (
	"net/http"

	"github.com/d782/tuut/bd"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "bd connection failed", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
