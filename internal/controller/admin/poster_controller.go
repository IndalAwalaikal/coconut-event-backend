package admin

import (
	"net/http"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/dto/response"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/service"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/util"
	"github.com/gorilla/mux"
)

type AdminPosterController struct {
    svc service.PosterService
}

func NewAdminPosterController(s service.PosterService) *AdminPosterController {
    return &AdminPosterController{svc: s}
}

func (c *AdminPosterController) Create(w http.ResponseWriter, r *http.Request) {
    title := r.FormValue("title")
    ptype := r.FormValue("type")
    date := r.FormValue("date")

    // try to get file; it may be optional
    file, header, err := r.FormFile("image")
    if err != nil {
        // missing file is allowed; set nil
        file = nil
        header = nil
    }
    if file != nil {
        defer file.Close()
    }

    p, err := c.svc.CreateFromForm(title, ptype, date, file, header)
    if err != nil {
        util.JSONError(w, 500, err.Error())
        return
    }
    util.JSON(w, 201, response.PosterFromModel(p))
}

func (c *AdminPosterController) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        util.JSONError(w, 400, "missing id")
        return
    }
    title := r.FormValue("title")
    ptype := r.FormValue("type")
    date := r.FormValue("date")
    file, header, err := r.FormFile("image")
    if err != nil {
        file = nil
        header = nil
    }
    if file != nil {
        defer file.Close()
    }
    p, err := c.svc.UpdateFromForm(id, title, ptype, date, file, header)
    if err != nil {
        util.JSONError(w, 500, err.Error())
        return
    }
    util.JSON(w, 200, response.PosterFromModel(p))
}

func (c *AdminPosterController) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        util.JSONError(w, 400, "missing id")
        return
    }
    if err := c.svc.Delete(id); err != nil {
        util.JSONError(w, 500, err.Error())
        return
    }
    util.JSON(w, 200, map[string]string{"status": "ok"})
}

func (c *AdminPosterController) List(w http.ResponseWriter, r *http.Request) {
    ps, err := c.svc.List()
    if err != nil {
        util.JSONError(w, 500, err.Error())
        return
    }
    util.JSON(w, 200, response.PostersFromModels(ps))
}
