package middleware

import (
	"log"
	"net/http"
	"time"
)

// RequestLogger logs the incoming HTTP requests with method, path, and duration
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

// Recovery recovers from panics and returns a 500 to the client
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Printf("PANIC: %v", rec)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
