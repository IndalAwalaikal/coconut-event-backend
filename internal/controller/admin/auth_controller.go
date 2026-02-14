package admin

import (
	"encoding/json"
	"net/http"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/service"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/util"
)

type AuthController struct {
    authService *service.AuthService
}

func NewAuthController(as *service.AuthService) *AuthController {
    return &AuthController{authService: as}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
    var req struct{ Username, Password string }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        util.JSONError(w, http.StatusBadRequest, "invalid payload")
        return
    }
    token, err := c.authService.Login(req.Username, req.Password)
    if err != nil {
        util.JSONError(w, http.StatusUnauthorized, "invalid credentials")
        return
    }
    util.JSON(w, http.StatusOK, map[string]string{"token": token})
}
