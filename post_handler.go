package main

import (
	"biblioV2/db"
	"database/sql"
	"encoding/json"
	"net/http"
)

func PostAbonneHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON data from the request body into an ABONNE struct
	var abonne db.ABONNE
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&abonne); err != nil {
		http.Error(w, "Failed to decode JSON data: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := db.PostAbonne(t, abonne)
	if err != nil {
		http.Error(w, "Failed to insert ABONNE data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If the insertion was successful, respond with a success status code
	w.WriteHeader(http.StatusCreated)
}

func PostCategorieAbonneHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON data from the request body into an CATEGORIE_ABONNE struct
	var categorie_abonne db.CATEGORIE_ABONNE
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&categorie_abonne); err != nil {
		http.Error(w, "Failed to decode JSON data: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := db.PostCategorieAbonne(t, categorie_abonne)
	if err != nil {
		http.Error(w, "Failed to insert CATEGORIE_ABONNE data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If the insertion was successful, respond with a success status code
	w.WriteHeader(http.StatusCreated)
}

func PostTypeOuvrageHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON data from the request body into an TYPE_OUVRAGE struct
	var type_ouvrage db.TYPE_OUVRAGE
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&type_ouvrage); err != nil {
		http.Error(w, "Failed to decode JSON data: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := db.PostTypeOuvrage(t, type_ouvrage)
	if err != nil {
		http.Error(w, "Failed to insert TYPE_OUVRAGE data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If the insertion was successful, respond with a success status code
	w.WriteHeader(http.StatusCreated)
}

func PostAuteurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON data from the request body into an AUTEUR struct
	var auteur db.AUTEUR
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&auteur); err != nil {
		http.Error(w, "Failed to decode JSON data: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := db.PostAuteur(t, auteur)
	if err != nil {
		http.Error(w, "Failed to insert AUTEUR data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If the insertion was successful, respond with a success status code
	w.WriteHeader(http.StatusCreated)
}

func PostEditeurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON data from the request body into an EDITEUR struct
	var editeur db.EDITEUR
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&editeur); err != nil {
		http.Error(w, "Failed to decode JSON data: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := db.PostEditeur(t, editeur)
	if err != nil {
		http.Error(w, "Failed to insert EDITEUR data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If the insertion was successful, respond with a success status code
	w.WriteHeader(http.StatusCreated)
}

func PostOuvrageHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON data from the request body into an OUVRAGE struct
	var ouvrage db.OUVRAGE
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ouvrage); err != nil {
		http.Error(w, "Failed to decode JSON data: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := db.PostOuvrage(t, ouvrage)
	if err != nil {
		http.Error(w, "Failed to insert OUVRAGE data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If the insertion was successful, respond with a success status code
	w.WriteHeader(http.StatusCreated)
}

func PostEcritHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON data from the request body into an ECRIT struct
	var ecrit db.ECRIT
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ecrit); err != nil {
		http.Error(w, "Failed to decode JSON data: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := db.PostEcrit(t, ecrit)
	if err != nil {
		http.Error(w, "Failed to insert ECRIT data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If the insertion was successful, respond with a success status code
	w.WriteHeader(http.StatusCreated)
}

func PostEmpruntHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON data from the request body into an EMPRUNT struct
	var emprunt db.EMPRUNTER
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&emprunt); err != nil {
		http.Error(w, "Failed to decode JSON data: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := db.PostEmprunter(t, emprunt)
	if err != nil {
		http.Error(w, "Failed to insert EMPRUNT data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If the insertion was successful, respond with a success status code
	w.WriteHeader(http.StatusCreated)
}
