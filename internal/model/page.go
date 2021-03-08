package model

import (
	"net/http"
	"strconv"
)

type Page struct {
	Page    int64
	PerPage int64
}

func ExtractPage(r *http.Request) Page {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	perPage, err := strconv.Atoi(r.URL.Query().Get("per_page"))
	if err != nil || perPage > 5 {
		perPage = 5
	}
	return Page{Page: int64(page), PerPage: int64(perPage)}
}
