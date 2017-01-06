package auth

import (
	"net/http"
)

func ShowMeYouToken(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
      return
    }
    next.ServeHTTP(w, r)
  })
}