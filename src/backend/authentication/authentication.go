package authentication

import (
	"log"
	"net/http"
	"os"
)

// Very simple access control list
func getTokenUsers() map[string]string {

	userToken := "05f717e5"
	userTokenFromEnv, ok := os.LookupEnv("APP_USER_TOKEN")
	if ok {
		userToken = userTokenFromEnv
	}

	return map[string]string{
		userToken: "authenticated user",
	}
}

// Middleware function, which will be called for each request
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-Auth-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Authentication middleware type
		tokenUsers := getTokenUsers()

		token := r.Header.Get("X-Auth-Token")

		if user, found := tokenUsers[token]; found {
			// We found the token in our map
			log.Printf("Authenticated user: %s\n", user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
