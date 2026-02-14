package middleware

import (
	"net/http"
	"strings"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/util"
)

func AdminAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        h := r.Header.Get("Authorization")
        if h == "" {
            util.JSONError(w, http.StatusUnauthorized, "missing authorization header")
            return
        }
        parts := strings.SplitN(h, " ", 2)
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            util.JSONError(w, http.StatusUnauthorized, "invalid authorization header")
            return
        }
        token := parts[1]
        claims, err := util.ParseAdminToken(token)
        if err != nil {
            util.JSONError(w, http.StatusUnauthorized, "invalid token")
            return
        }
        // set claims in context if desired (omitted for brevity)
        _ = claims
        next.ServeHTTP(w, r)
    })
}
