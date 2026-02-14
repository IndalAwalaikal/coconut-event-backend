package public

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/dto/request"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/service"
)

type RegistrationController struct {
    Service *service.RegistrationService
}

func NewRegistrationController(s *service.RegistrationService) *RegistrationController {
    return &RegistrationController{Service: s}
}

// POST /api/registrations
// Accepts multipart/form-data with fields: event_id, name, whatsapp, institution, and optional file under `proof`.
func (c *RegistrationController) Create(w http.ResponseWriter, r *http.Request) {
    // limit request size to 8MB + file
    if err := r.ParseMultipartForm(10 << 20); err != nil && err != http.ErrNotMultipart {
        http.Error(w, "invalid form data", http.StatusBadRequest)
        return
    }

    // check for multipart file
    var header *multipart.FileHeader
    if r.MultipartForm != nil {
        fhs := r.MultipartForm.File["proof"]
        if len(fhs) > 0 {
            header = fhs[0]
        }
    }

    // fallback to urlencoded forms
    req := request.CreateRegistrationRequest{
        EventID: r.FormValue("event_id"),
        Name: r.FormValue("name"),
        Whatsapp: r.FormValue("whatsapp"),
        Institution: r.FormValue("institution"),
        FileName: r.FormValue("file_name"),
    }

    // create registration
    var f multipart.File
    var fh *multipart.FileHeader
    if header != nil {
        f, _ = header.Open()
        fh = header
        defer func() {
            if f != nil {
                f.Close()
            }
        }()
    }

    reg, err := c.Service.Register(req.EventID, req.Name, req.Whatsapp, req.Institution, f, fh, req.FileName)
    if err != nil {
        log.Printf("registration error: %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    _ = json.NewEncoder(w).Encode(reg)
}
