package util

import (
	"net/http"
	"strconv"
)

const (
    DefaultPage = 1
    DefaultLimit = 20
    MaxLimit = 200
)

// ParsePagination reads ?page and ?limit from request and returns limit, offset, page
func ParsePagination(r *http.Request) (limit, offset, page int) {
    page = DefaultPage
    limit = DefaultLimit
    if p := r.URL.Query().Get("page"); p != "" {
        if pv, err := strconv.Atoi(p); err == nil && pv > 0 { page = pv }
    }
    if l := r.URL.Query().Get("limit"); l != "" {
        if lv, err := strconv.Atoi(l); err == nil && lv > 0 {
            if lv > MaxLimit { lv = MaxLimit }
            limit = lv
        }
    }
    offset = (page - 1) * limit
    return
}
