package apiServer

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
)

func LogMidellware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, duration)
	}
}

func RecoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
				http.Error(w, "Internal Server Error", 500)
			}
		}()
		next(w, r)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		id, err := strconv.Atoi(cookie.Value)
		if err != nil {
			log.Printf("Cookie not found %v", err)
		}
		ctx := context.WithValue(r.Context(), "KeyUserID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
