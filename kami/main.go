package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/guregu/kami"
	"golang.org/x/net/context"
)

func main() {
	kami.Get("/contacts", getContacts)
	kami.Serve()
}

func getContacts(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
) {
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
		}
	}()

	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		page = 1
	}

	perPage, err := strconv.Atoi(r.FormValue("per_page"))
	if err != nil {
		perPage = 100
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err = json.NewEncoder(w).Encode(
		NewContactQuery(page, perPage).All())

	if err != nil {
		log.Print(err)
	}
}
