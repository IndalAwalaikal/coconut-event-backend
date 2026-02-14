package admin

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type DashboardController struct {
    DB *sql.DB
}

func NewDashboardController(db *sql.DB) *DashboardController {
    return &DashboardController{DB: db}
}

func (c *DashboardController) Get(w http.ResponseWriter, r *http.Request) {
    var out struct {
        TotalEvents     int `json:"totalEvents"`
        ActiveEvents    int `json:"activeEvents"`
        TotalRegistrants int `json:"totalRegistrants"`
    }
    // total events
    if err := c.DB.QueryRow(`SELECT COUNT(*) FROM events`).Scan(&out.TotalEvents); err != nil {
        log.Printf("dashboard total events: %v", err)
    }
    if err := c.DB.QueryRow(`SELECT COUNT(*) FROM events WHERE available = 1`).Scan(&out.ActiveEvents); err != nil {
        log.Printf("dashboard active events: %v", err)
    }
    if err := c.DB.QueryRow(`SELECT COUNT(*) FROM registrations`).Scan(&out.TotalRegistrants); err != nil {
        log.Printf("dashboard total regs: %v", err)
    }
    b, _ := json.Marshal(out)
    w.Header().Set("Content-Type", "application/json")
    w.Write(b)
}
