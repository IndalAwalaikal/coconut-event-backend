package middleware

import (
	"net/http"
	"os"
)

// CORS adds permissive CORS headers for development. In production set
// CORS_ALLOWED and tighten origins.
func CORS(next http.Handler) http.Handler {
    allowed := os.Getenv("CORS_ALLOWED")
    if allowed == "" {
        allowed = "*"
    }
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", allowed)
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        next.ServeHTTP(w, r)
    })
}
