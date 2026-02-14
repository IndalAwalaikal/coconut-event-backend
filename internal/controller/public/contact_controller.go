package public

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/util"
)

type contactPayload struct {
    Name    string `json:"name"`
    Email   string `json:"email"`
    Message string `json:"message"`
}

type ContactController struct{}

func NewContactController() *ContactController { return &ContactController{} }

// Create handles POST /api/contact
func (c *ContactController) Create(w http.ResponseWriter, r *http.Request) {
    var p contactPayload
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        util.JSONError(w, http.StatusBadRequest, "invalid payload")
        return
    }
    if p.Name == "" || p.Email == "" || p.Message == "" {
        util.JSONError(w, http.StatusBadRequest, "name,email,message required")
        return
    }
    // persist to file for now
    dir := "storage/contacts"
    _ = os.MkdirAll(dir, 0o755)
    fname := filepath.Join(dir, time.Now().Format("20060102-150405")+".json")
    b, _ := json.MarshalIndent(p, "", "  ")
    if err := ioutil.WriteFile(fname, b, 0o644); err != nil {
        util.JSONError(w, http.StatusInternalServerError, "failed to save")
        return
    }
    util.JSON(w, http.StatusCreated, map[string]string{"status": "ok"})
}
