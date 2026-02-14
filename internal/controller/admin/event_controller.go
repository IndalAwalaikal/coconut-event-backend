package admin

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/service"
)

type AdminEventController struct {
    Service *service.EventService
}

func NewAdminEventController(s *service.EventService) *AdminEventController {
    return &AdminEventController{Service: s}
}

// POST /api/admin/events
// Accepts multipart/form-data for poster
func (c *AdminEventController) Create(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseMultipartForm(10 << 20); err != nil && err != http.ErrNotMultipart {
        http.Error(w, "invalid form", http.StatusBadRequest)
        return
    }
    // collect form values into map
    m := map[string]interface{}{}
    for k, v := range r.Form {
        if len(v) == 1 {
            m[k] = v[0]
        } else {
            // multiple values -> keep as []string
            m[k] = v
        }
    }
    var file multipart.File
    var header *multipart.FileHeader
    if r.MultipartForm != nil {
        fhs := r.MultipartForm.File["poster"]
        if len(fhs) > 0 {
            header = fhs[0]
            f, _ := header.Open()
            file = f
            defer f.Close()
        }
    }
    ev, err := c.Service.Create(m, file, header)
    if err != nil {
        log.Printf("create event err: %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    _ = json.NewEncoder(w).Encode(ev)
}

// PUT /api/admin/events/{id}
func (c *AdminEventController) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        http.Error(w, "id required", http.StatusBadRequest)
        return
    }
    if err := r.ParseMultipartForm(10 << 20); err != nil && err != http.ErrNotMultipart {
        http.Error(w, "invalid form", http.StatusBadRequest)
        return
    }
    m := map[string]interface{}{}
    for k, v := range r.Form {
        if len(v) == 1 { m[k] = v[0] } else { m[k] = v }
    }
    var file multipart.File
    var header *multipart.FileHeader
    if r.MultipartForm != nil {
        fhs := r.MultipartForm.File["poster"]
        if len(fhs) > 0 {
            header = fhs[0]
            f, _ := header.Open()
            file = f
            defer f.Close()
        }
    }
    ev, err := c.Service.Update(id, m, file, header)
    if err != nil {
        log.Printf("update event err: %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(ev)
}

// DELETE /api/admin/events/{id}
func (c *AdminEventController) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        http.Error(w, "id required", http.StatusBadRequest)
        return
    }
    if err := c.Service.Delete(id); err != nil {
        log.Printf("delete event err: %v", err)
        http.Error(w, "failed", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
