package main

import (
	"biblioV2/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func GetAbonneByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	abonne, err := db.GetAbonneById(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(abonne)
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

func GetCategorieAbonneByIdHandler(w http.ResponseWriter, r *http.Request) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer t.Close()

	//get id from url
	idStr := mux.Vars(r)["id"]

	categorie_abonne, err := db.GetCategorieAbonneById(t, idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(categorie_abonne)
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

func GetTypeOuvrageByIdHandler(w http.ResponseWriter, r *http.Request) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer t.Close()

	//get id from url
	idStr := mux.Vars(r)["id"]

	type_ouvrage, err := db.GetTypeOuvrageById(t, idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(type_ouvrage)
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

func GetAuteurByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	auteur, err := db.GetAuteurById(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(auteur)
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

func GetEditeurByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	editeur, err := db.GetEditeurById(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(editeur)
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

func GetOuvrageByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	ouvrage, err := db.GetOuvrageById(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(ouvrage)
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

func GetEcritByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	ecrit, err := db.GetEcritById(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(ecrit)
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

func GetEmpruntByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	emprunter, err := db.GetEmprunterById(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(emprunter)
}
