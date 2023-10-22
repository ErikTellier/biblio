package main

import (
	"biblioV2/db"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func RetourEmprunterHandler(w http.ResponseWriter, r *http.Request) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer t.Close()

	//get id from url
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//update emprunter
	err = db.RetourEmprunter(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
