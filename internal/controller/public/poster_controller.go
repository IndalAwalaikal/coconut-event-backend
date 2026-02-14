package public

import (
	"net/http"

	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/dto/response"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/service"
	"github.com/IndalAwalaikal/coconut-event-hub/backend/internal/util"
	"github.com/gorilla/mux"
)

type PosterController struct {
    svc service.PosterService
}

func NewPosterController(s service.PosterService) *PosterController {
    return &PosterController{svc: s}
}

func (c *PosterController) List(w http.ResponseWriter, r *http.Request) {
    ps, err := c.svc.List()
    if err != nil {
        util.JSONError(w, 500, err.Error())
        return
    }
    util.JSON(w, 200, response.PostersFromModels(ps))
}

func (c *PosterController) Get(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        util.JSONError(w, 400, "missing id")
        return
    }
    p, err := c.svc.GetByID(id)
    if err != nil {
        util.JSONError(w, 500, err.Error())
        return
    }
    util.JSON(w, 200, response.PosterFromModel(p))
}
