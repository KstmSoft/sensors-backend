package controllers

import (
	"bytes"
	"io"
	"net/http"
	"sensors/app/services"

	"github.com/gorilla/mux"
)

func GetLogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		b := bytes.NewBufferString(services.ReadLogs(id))
		w.Header().Set("Content-Disposition", "attachment; filename="+id+".csv")
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		io.Copy(w, b)
	}
}
