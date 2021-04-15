package middleware

import (
	"net/http"

	"github.com/FourLineCode/financer/pkg/handler"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement jwt
		if token := r.Header.Get("auth"); token != "token" {
			handler.ResponseError(w, http.StatusForbidden, "Authorization failed!")
			return
		}

		next(w, r)
	})
}
