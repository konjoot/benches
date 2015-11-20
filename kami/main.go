package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/guregu/kami"

	// _ "net/http/pprof"
)

func main() {
	// go func() {
	// 	http.ListenAndServe("localhost:6060", nil)
	// }()
	kami.Get("/contacts", getContacts)
	kami.Serve()
}

func getContacts(
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

	// NewContactQuery(page, perPage).All().JSON(w)
	out := NewContactQuery(page, perPage).All()
	err = json.NewEncoder(w).Encode(out)

	if err != nil {
		log.Print(err)
	}
}
