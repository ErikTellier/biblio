package main

import (
	"biblioV2/db"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func PutAbonneHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {
	//get id from url
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get abonne from body
	var abonne db.ABONNE
	err = json.NewDecoder(r.Body).Decode(&abonne)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	abonne.ID_ABONNE = id

	//update abonne
	err = db.PutAbonne(t, abonne)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PutAuteurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	//get id from url
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get auteur from body
	var auteur db.AUTEUR
	err = json.NewDecoder(r.Body).Decode(&auteur)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auteur.ID_AUTEUR = id

	//update auteur
	err = db.PutAuteur(t, auteur)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PutEditeurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	//get id from url
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get editeur from body
	var editeur db.EDITEUR
	err = json.NewDecoder(r.Body).Decode(&editeur)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	editeur.ID_EDITEUR = id

	//update editeur
	err = db.PutEditeur(t, editeur)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PutOuvrageHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	//get id from url
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get ouvrage from body
	var ouvrage db.OUVRAGE
	err = json.NewDecoder(r.Body).Decode(&ouvrage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ouvrage.ID_OUVRAGE = id

	//update ouvrage
	err = db.PutOuvrage(t, ouvrage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PutEcritHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	//get id from url
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get ecrit from body
	var ecrit db.ECRIT
	err = json.NewDecoder(r.Body).Decode(&ecrit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ecrit.ID_ECRIT = id

	//update ecrit
	err = db.PutEcrit(t, ecrit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PutEmpruntHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	//get id from url
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get emprunter from body
	var emprunter db.EMPRUNTER
	err = json.NewDecoder(r.Body).Decode(&emprunter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	emprunter.ID_EMPRUNT = id

	//update emprunter
	err = db.PutEmprunter(t, emprunter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
