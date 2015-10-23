package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/guregu/kami"
	"golang.org/x/net/context"
)

func main() {
	kami.Get("/contacts/", getContacts)
	kami.Serve()
}

func getContacts(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		page = 1
	}

	perPage, err := strconv.Atoi(r.FormValue("perPage"))
	if err != nil {
		perPage = 100
	}

	json.NewEncoder(w).Encode(
		NewContactQuery(page, perPage).All())
}
