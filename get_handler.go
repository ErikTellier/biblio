package main

import (
	"biblioV2/db"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetAbonneHandler
// @Summary Get a list of abonne
// @Description Get a list of abonne
// @ID get-abonne
// @Produce json
// @Success 200
// @Router /api/abonne [get]
func GetAbonneHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {
	abonnes, err := db.GetAllAbonnes(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abonnes)
}

// GetAbonneByIdHandler
// @Summary Récupérer un abonné par ID
// @Description Récupère un abonné en fonction de son ID
// @ID get-abonne-by-id
// @Produce json
// @Param id path int true "ID de l'abonné à récupérer"
// @Success 200
// @Router /api/abonne/{id} [get]
func GetAbonneByIdHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {
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

// GetCategorieAbonneHandler
// @Summary Récupérer la liste des catégories d'abonnés
// @Description Récupère la liste complète des catégories d'abonnés
// @ID get-categorie-abonne
// @Produce json
// @Success 200
// @Router /api/categorie_abonne [get]
func GetCategorieAbonneHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	categorie_abonnes, err := db.GetAllCategorieAbonnes(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categorie_abonnes)
}

// GetCategorieAbonneByIdHandler
// @Summary Récupérer une catégorie d'abonné par ID
// @Description Récupère une catégorie d'abonné en fonction de son ID
// @ID get-categorie-abonne-by-id
// @Produce json
// @Param id path int true "ID de la catégorie d'abonné à récupérer"
// @Success 200
// @Router /api/categorie_abonne/{id} [get]
func GetCategorieAbonneByIdHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	idStr := mux.Vars(r)["id"]

	categorie_abonne, err := db.GetCategorieAbonneById(t, idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(categorie_abonne)
}

// GetTypeOuvrageHandler
// @Summary Récupérer la liste des types d'ouvrages
// @Description Récupère la liste complète des types d'ouvrages
// @ID get-type-ouvrage
// @Produce json
// @Success 200
// @Router /api/type_ouvrage [get]
func GetTypeOuvrageHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	type_ouvrages, err := db.GetAllTypeOuvrage(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(type_ouvrages)
}

// GetTypeOuvrageByIdHandler
// @Summary Récupérer un type d'ouvrage par ID
// @Description Récupère un type d'ouvrage en fonction de son ID
// @ID get-type-ouvrage-by-id
// @Produce json
// @Param id path int true "ID du type d'ouvrage à récupérer"
// @Success 200
// @Router /api/type_ouvrage/{id} [get]
func GetTypeOuvrageByIdHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	idStr := mux.Vars(r)["id"]

	type_ouvrage, err := db.GetTypeOuvrageById(t, idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(type_ouvrage)
}

// GetAuteurHandler
// @Summary Récupérer la liste des auteurs
// @Description Récupère la liste complète des auteurs
// @ID get-auteur
// @Produce json
// @Success 200
// @Router /api/auteur [get]
func GetAuteurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	auteurs, err := db.GetAllAuteur(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auteurs)
}

// GetAuteurByIdHandler
// @Summary Récupérer un auteur par ID
// @Description Récupère un auteur en fonction de son ID
// @ID get-auteur-by-id
// @Produce json
// @Param id path int true "ID de l'auteur à récupérer"
// @Success 200
// @Router /api/auteur/{id} [get]
func GetAuteurByIdHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

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

// GetEditeurHandler
// @Summary Récupérer la liste des éditeurs
// @Description Récupère la liste complète des éditeurs
// @ID get-editeur
// @Produce json
// @Success 200
// @Router /api/editeur [get]
func GetEditeurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	editeurs, err := db.GetAllEditeur(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(editeurs)
}

// GetEditeurByIdHandler
// @Summary Récupérer un éditeur par ID
// @Description Récupère un éditeur en fonction de son ID
// @ID get-editeur-by-id
// @Produce json
// @Param id path int true "ID de l'éditeur à récupérer"
// @Success 200
// @Router /api/editeur/{id} [get]
func GetEditeurByIdHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

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

// GetOuvrageHandler
// @Summary Récupérer la liste des ouvrages
// @Description Récupère la liste complète des ouvrages
// @ID get-ouvrage
// @Produce json
// @Success 200
// @Router /api/ouvrage [get]
func GetOuvrageHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	ouvrages, err := db.GetAllOuvrage(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ouvrages)
}

// GetOuvrageByIdHandler
// @Summary Récupérer un ouvrage par ID
// @Description Récupère un ouvrage en fonction de son ID
// @ID get-ouvrage-by-id
// @Produce json
// @Param id path int true "ID de l'ouvrage à récupérer"
// @Success 200
// @Router /api/ouvrage/{id} [get]
func GetOuvrageByIdHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

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

// GetEcritHandler
// @Summary Récupérer la liste des écrits
// @Description Récupère la liste complète des écrits
// @ID get-ecrit
// @Produce json
// @Success 200
// @Router /api/ecrit [get]
func GetEcritHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	ecrits, err := db.GetAllEcrit(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ecrits)
}

// GetEcritByIdHandler
// @Summary Récupérer un écrit par ID
// @Description Récupère un écrit en fonction de son ID
// @ID get-ecrit-by-id
// @Produce json
// @Param id path int true "ID de l'écrit à récupérer"
// @Success 200
// @Router /api/ecrit/{id} [get]
func GetEcritByIdHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

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

// GetEmprunterHandler
// @Summary Récupérer la liste des emprunteurs
// @Description Récupère la liste complète des emprunteurs
// @ID get-emprunter
// @Produce json
// @Success 200
// @Router /api/emprunter [get]
func GetEmprunterHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	emprunters, err := db.GetAllEmprunter(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emprunters)
}

// GetEmpruntByIdHandler
// @Summary Récupérer un emprunt par ID
// @Description Récupère un emprunt en fonction de son ID
// @ID get-emprunt-by-id
// @Produce json
// @Param id path int true "ID de l'emprunt à récupérer"
// @Success 200
// @Router /api/emprunter/{id} [get]
func GetEmpruntByIdHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {
	t, err := db.OpenDBConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer t.Close()

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
