package main

import (
	"biblioV2/db"
	"encoding/json"
	"net/http"
)

func GetAbonneHandler(w http.ResponseWriter, r *http.Request) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer t.Close()

	abonnes, err := db.GetAllAbonnes(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abonnes)
}

func GetCategorieAbonneHandler(w http.ResponseWriter, r *http.Request) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer t.Close()

	categorie_abonnes, err := db.GetAllCategorieAbonnes(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categorie_abonnes)
}

func GetTypeOuvrageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer t.Close()

	type_ouvrages, err := db.GetAllTypeOuvrage(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(type_ouvrages)
}

func GetAuteurHandler(w http.ResponseWriter, r *http.Request) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer t.Close()

	auteurs, err := db.GetAllAuteur(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auteurs)
}

func GetEditeurHandler(w http.ResponseWriter, r *http.Request) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer t.Close()

	editeurs, err := db.GetAllEditeur(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(editeurs)
}

func GetOuvrageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer t.Close()

	ouvrages, err := db.GetAllOuvrage(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ouvrages)
}

func GetEcritHandler(w http.ResponseWriter, r *http.Request) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer t.Close()

	ecrits, err := db.GetAllEcrit(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ecrits)
}

func GetEmprunterHandler(w http.ResponseWriter, r *http.Request) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer t.Close()

	emprunters, err := db.GetAllEmprunter(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emprunters)
}
