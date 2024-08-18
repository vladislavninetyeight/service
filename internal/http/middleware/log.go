package middleware

import (
	"log"
	"net/http"
	"time"
)

func RequestLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()
			next.ServeHTTP(w, r)
			log.Print(time.Since(t))
		}
		return http.HandlerFunc(fn)
	}
}
