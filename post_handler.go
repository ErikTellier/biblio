package main

import (
	"biblioV2/db"
	"database/sql"
	"encoding/json"
	"net/http"
)

// PostAbonneHandler
// @Summary Créer un nouvel abonné
// @Description Crée un nouvel abonné en utilisant les données JSON fournies dans la requête POST
// @ID create-abonne
// @Produce json
// @Consume json
// @Success 201 "Abonné créé avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de l'insertion des données ABONNE"
// @Router /api/abonne [post]
func PostAbonneHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

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

	w.WriteHeader(http.StatusCreated)
}

// PostCategorieAbonneHandler
// @Summary Créer une nouvelle catégorie d'abonné
// @Description Crée une nouvelle catégorie d'abonné en utilisant les données JSON fournies dans la requête POST
// @ID create-categorie-abonne
// @Produce json
// @Consume json
// @Success 201 "Catégorie d'abonné créée avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de l'insertion des données CATEGORIE_ABONNE"
// @Router /api/categorie_abonne [post]
func PostCategorieAbonneHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

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

	w.WriteHeader(http.StatusCreated)
}

// PostTypeOuvrageHandler
// @Summary Créer un nouveau type d'ouvrage
// @Description Crée un nouveau type d'ouvrage en utilisant les données JSON fournies dans la requête POST
// @ID create-type-ouvrage
// @Produce json
// @Consume json
// @Success 201 "Type d'ouvrage créé avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de l'insertion des données TYPE_OUVRAGE"
// @Router /api/type_ouvrage [post]
func PostTypeOuvrageHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

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

	w.WriteHeader(http.StatusCreated)
}

// PostAuteurHandler
// @Summary Créer un nouvel auteur
// @Description Crée un nouvel auteur en utilisant les données JSON fournies dans la requête POST
// @ID create-auteur
// @Produce json
// @Consume json
// @Success 201 "Auteur créé avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de l'insertion des données AUTEUR"
// @Router /api/auteur [post]
func PostAuteurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

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

	w.WriteHeader(http.StatusCreated)
}

// PostEditeurHandler
// @Summary Créer un nouvel éditeur
// @Description Crée un nouvel éditeur en utilisant les données JSON fournies dans la requête POST
// @ID create-editeur
// @Produce json
// @Consume json
// @Success 201 "Éditeur créé avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de l'insertion des données EDITEUR"
// @Router /api/editeur [post]
func PostEditeurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

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

	w.WriteHeader(http.StatusCreated)
}

// PostOuvrageHandler
// @Summary Créer un nouvel ouvrage
// @Description Crée un nouvel ouvrage en utilisant les données JSON fournies dans la requête POST
// @ID create-ouvrage
// @Produce json
// @Consume json
// @Success 201 "Ouvrage créé avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de l'insertion des données OUVRAGE"
// @Router /api/ouvrage [post]
func PostOuvrageHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

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

	w.WriteHeader(http.StatusCreated)
}

// PostEcritHandler
// @Summary Créer un nouvel écrit
// @Description Crée un nouvel écrit en utilisant les données JSON fournies dans la requête POST
// @ID create-ecrit
// @Produce json
// @Consume json
// @Success 201 "Écrit créé avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de l'insertion des données ECRIT"
// @Router /api/ecrit [post]
func PostEcritHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

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

	w.WriteHeader(http.StatusCreated)
}

// PostEmpruntHandler
// @Summary Créer un nouvel emprunt
// @Description Crée un nouvel emprunt en utilisant les données JSON fournies dans la requête POST
// @ID create-emprunt
// @Produce json
// @Consume json
// @Success 201 "Emprunt créé avec succès"
// @Failure 400 "Échec de décodage des données JSON"
// @Failure 500 "Échec de l'insertion des données EMPRUNT"
// @Router /api/emprunt [post]
func PostEmpruntHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

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

	w.WriteHeader(http.StatusCreated)
}
