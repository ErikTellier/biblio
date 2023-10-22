package main

import (
	"biblioV2/db"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// PutAbonneHandler
// @Summary Mettre à jour un abonné existant
// @Description Met à jour un abonné existant en utilisant les données JSON fournies dans la requête PUT
// @ID update-abonne
// @Produce json
// @Consume json
// @Param id path int true "ID de l'abonné à mettre à jour"
// @Success 200 "Abonné mis à jour avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de la mise à jour de l'abonné"
// @Router /api/abonne/{id} [put]
func PutAbonneHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var abonne db.ABONNE
	err = json.NewDecoder(r.Body).Decode(&abonne)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	abonne.ID_ABONNE = id

	err = db.PutAbonne(t, abonne)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// PutAuteurHandler
// @Summary Mettre à jour un auteur existant
// @Description Met à jour un auteur existant en utilisant les données JSON fournies dans la requête PUT
// @ID update-auteur
// @Produce json
// @Consume json
// @Param id path int true "ID de l'auteur à mettre à jour"
// @Success 200 "Auteur mis à jour avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de la mise à jour de l'auteur"
// @Router /api/auteur/{id} [put]
func PutAuteurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var auteur db.AUTEUR
	err = json.NewDecoder(r.Body).Decode(&auteur)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auteur.ID_AUTEUR = id

	err = db.PutAuteur(t, auteur)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// PutEditeurHandler
// @Summary Mettre à jour un éditeur existant
// @Description Met à jour un éditeur existant en utilisant les données JSON fournies dans la requête PUT
// @ID update-editeur
// @Produce json
// @Consume json
// @Param id path int true "ID de l'éditeur à mettre à jour"
// @Success 200 "Éditeur mis à jour avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de la mise à jour de l'éditeur"
// @Router /api/editeur/{id} [put]
func PutEditeurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var editeur db.EDITEUR
	err = json.NewDecoder(r.Body).Decode(&editeur)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	editeur.ID_EDITEUR = id

	err = db.PutEditeur(t, editeur)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// PutOuvrageHandler
// @Summary Mettre à jour un ouvrage existant
// @Description Met à jour un ouvrage existant en utilisant les données JSON fournies dans la requête PUT
// @ID update-ouvrage
// @Produce json
// @Consume json
// @Param id path int true "ID de l'ouvrage à mettre à jour"
// @Success 200 "Ouvrage mis à jour avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de la mise à jour de l'ouvrage"
// @Router /api/ouvrage/{id} [put]
func PutOuvrageHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var ouvrage db.OUVRAGE
	err = json.NewDecoder(r.Body).Decode(&ouvrage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ouvrage.ID_OUVRAGE = id

	err = db.PutOuvrage(t, ouvrage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// PutEcritHandler
// @Summary Mettre à jour un écrit existant
// @Description Met à jour un écrit existant en utilisant les données JSON fournies dans la requête PUT
// @ID update-ecrit
// @Produce json
// @Consume json
// @Param id path int true "ID de l'écrit à mettre à jour"
// @Success 200 "Écrit mis à jour avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de la mise à jour de l'écrit"
// @Router /api/ecrit/{id} [put]
func PutEcritHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var ecrit db.ECRIT
	err = json.NewDecoder(r.Body).Decode(&ecrit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ecrit.ID_ECRIT = id

	err = db.PutEcrit(t, ecrit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// PutEmpruntHandler
// @Summary Mettre à jour un emprunt existant
// @Description Met à jour un emprunt existant en utilisant les données JSON fournies dans la requête PUT
// @ID update-emprunt
// @Produce json
// @Consume json
// @Param id path int true "ID de l'emprunt à mettre à jour"
// @Success 200 "Emprunt mis à jour avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de la mise à jour de l'emprunt"
// @Router /api/emprunt/{id} [put]
func PutEmpruntHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var emprunter db.EMPRUNTER
	err = json.NewDecoder(r.Body).Decode(&emprunter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	emprunter.ID_EMPRUNT = id

	err = db.PutEmprunter(t, emprunter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
