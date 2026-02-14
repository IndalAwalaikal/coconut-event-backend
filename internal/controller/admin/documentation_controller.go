package admin

import (
	"net/http"
    "fmt"
	"strconv"

	respdto "github.com/IndalAwalaikal/coconut-event-hub/backend/internal/dto/response"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/service"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/util"
	"github.com/gorilla/mux"
)

type AdminDocumentationController struct {
    svc *service.DocumentationService
}

func NewAdminDocumentationController(svc *service.DocumentationService) *AdminDocumentationController {
    return &AdminDocumentationController{svc: svc}
}

// Create handles POST /api/admin/documentations
// expects multipart/form-data with fields: category, category_label, event_title, year, event_id (optional), description and files images[]
func (c *AdminDocumentationController) Create(w http.ResponseWriter, r *http.Request) {
    doc, err := c.svc.CreateFromForm(r)
    if err != nil {
         fmt.Printf("❌ ERROR: %v\n", err)
        util.JSONError(w, http.StatusBadRequest, err.Error())
        return
    }
    // map to response DTO
    out := respdto.DocumentationFromModel(doc)
    util.JSON(w, http.StatusCreated, out)
}

// Update handles PUT /api/admin/documentations/{id}
func (c *AdminDocumentationController) Update(w http.ResponseWriter, r *http.Request) {
    // id from path
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        util.JSONError(w, http.StatusBadRequest, "missing id")
        return
    }
    doc, err := c.svc.UpdateFromForm(id, r)
    if err != nil {
        util.JSONError(w, http.StatusBadRequest, err.Error())
        return
    }
    out := respdto.DocumentationFromModel(doc)
    util.JSON(w, http.StatusOK, out)
}

// Delete handles DELETE /api/admin/documentations/{id}
func (c *AdminDocumentationController) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        util.JSONError(w, http.StatusBadRequest, "missing id")
        return
    }
    if err := c.svc.Delete(id); err != nil {
        util.JSONError(w, http.StatusInternalServerError, "failed to delete")
        return
    }
    util.JSON(w, http.StatusNoContent, nil)
}

func (c *AdminDocumentationController) List(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    cat := r.URL.Query().Get("category")
    yearStr := r.URL.Query().Get("year")
    year := 0
    if yearStr != "" {
        if y, err := strconv.Atoi(yearStr); err == nil {
            year = y
        }
    }
    docs, err := c.svc.List(cat, year, q, 0, 0)
    if err != nil {
        util.JSONError(w, http.StatusInternalServerError, "Gagal memuat dokumentasi")
        return
    }
    // map to response DTOs
    out := respdto.DocumentationsFromModels(docs)
    util.JSON(w, http.StatusOK, out)
}