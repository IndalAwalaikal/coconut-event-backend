package middleware

import (
	"net/http"
	"os"
)

// CORS adds permissive CORS headers for development. In production set
// CORS_ALLOWED and tighten origins.
func CORS(next http.Handler) http.Handler {
    allowed := os.Getenv("CORS_ALLOWED")

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        origin := r.Header.Get("Origin")

        if origin == allowed {
            w.Header().Set("Access-Control-Allow-Origin", origin)
            w.Header().Set("Access-Control-Allow-Credentials", "true")
        }

        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
}
