package admin

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/service"
)

type AdminRegistrationController struct {
    Service *service.RegistrationService
    ExportService *service.ExportService
}

func NewAdminRegistrationController(s *service.RegistrationService) *AdminRegistrationController {
    return &AdminRegistrationController{Service: s, ExportService: service.NewExportService(s.Repo)}
}

// GET /api/admin/registrations?event_id={id}
func (c *AdminRegistrationController) List(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("event_id")
    if q == "" {
        http.Error(w, "event_id required", http.StatusBadRequest)
        return
    }
    regs, err := c.Service.ListRegistrationsByEvent(q)
    if err != nil {
        log.Printf("admin list regs err: %v", err)
        http.Error(w, "internal", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(regs)
}

// GET /api/admin/registrations/{id}
func (c *AdminRegistrationController) Get(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        http.Error(w, "id required", http.StatusBadRequest)
        return
    }
    reg, err := c.Service.GetRegistration(id)
    if err != nil {
        log.Printf("admin get reg err: %v", err)
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(reg)
}

// GET /api/admin/registrations/export?event_id={id}
func (c *AdminRegistrationController) ExportCSV(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("event_id")
    if q == "" {
        http.Error(w, "event_id required", http.StatusBadRequest)
        return
    }
    data, err := c.ExportService.ExportRegistrationsCSV(q)
    if err != nil {
        log.Printf("export err: %v", err)
        http.Error(w, "failed to export", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Disposition", "attachment; filename=registrants_"+q+".csv")
    w.Header().Set("Content-Type", "text/csv")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
}
