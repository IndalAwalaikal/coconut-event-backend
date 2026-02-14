package public

import (
	"log"
	"net/http"
	"strconv"

	respdto "github.com/IndalAwalaikal/coconut-event-hub/backend/internal/dto/response"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/service"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/util"
	"github.com/gorilla/mux"
)

type DocumentationController struct {
    svc *service.DocumentationService
}

func NewDocumentationController(svc *service.DocumentationService) *DocumentationController {
    return &DocumentationController{svc: svc}
}

// List handles GET /api/documentations
func (c *DocumentationController) List(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    cat := r.URL.Query().Get("category")
    yearStr := r.URL.Query().Get("year")
    year := 0
    if yearStr != "" {
        if y, err := strconv.Atoi(yearStr); err == nil { year = y }
    }
    docs, err := c.svc.List(cat, year, q, 0, 0)
    if err != nil {
        util.JSONError(w, http.StatusInternalServerError, "failed to list documentations")
        return
    }
    // debug log count and current DB
    dbname := ""
    if dbn, err := c.svc.DebugDatabase(); err == nil {
        dbname = dbn
    } else {
        dbname = "<unknown>"
    }
    log.Printf("public documentations: found=%d category=%q year=%d q=%q db=%s", len(docs), cat, year, q, dbname)
    out := respdto.DocumentationsFromModels(docs)
    util.JSON(w, http.StatusOK, out)
}

// Get handles GET /api/documentations/{id}
func (c *DocumentationController) Get(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        util.JSONError(w, http.StatusBadRequest, "missing id")
        return
    }
    doc, err := c.svc.GetByID(id)
    if err != nil {
        util.JSONError(w, http.StatusNotFound, "not found")
        return
    }
    out := respdto.DocumentationFromModel(doc)
    util.JSON(w, http.StatusOK, out)
}
