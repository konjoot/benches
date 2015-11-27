package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/guregu/kami"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func init() {
	DBConn()
	// PgxDBConn()
	// PgDBConn()
}

func main() {
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

	contacts := NewContactQuery(page, perPage).All()
	// contacts := NewPGXContactQuery(page, perPage).All()
	// contacts := NewPGContactQuery(page, perPage).All()
	buf := bufPool.Get().(*bytes.Buffer)
	err = json.NewEncoder(buf).Encode(contacts)
	w.Write(buf.Bytes())
	buf.Reset()
	bufPool.Put(buf)

	if err != nil {
		log.Print(err)
	}
}
