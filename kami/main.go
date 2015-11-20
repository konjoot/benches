package main

import (
	// _ "net/http/pprof"
	// "encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/guregu/kami"
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

	data, err := NewContactQuery(page, perPage).JSON()
	// err = json.NewEncoder(w).Encode(
	// 	NewContactQuery(page, perPage).All())
	if err != nil {
		log.Print(err)
	}

	w.Write(data)
}
