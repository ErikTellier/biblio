package main

import (
	"biblioV2/db"
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// DeleteAbonneHandler
// @Summary Supprimer un abonné par ID
// @Description Supprime un abonné en fonction de son ID
// @ID delete-abonne
// @Produce json
// @Param id path int true "ID de l'abonné à supprimer"
// @Success 204 "Abonné supprimé avec succès"
// @Router /api/abonne/{id} [delete]
func DeleteAbonneHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.DeleteAbonne(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteCategorieAbonneHandler
// @Summary Supprimer une catégorie d'abonné par ID
// @Description Supprime une catégorie d'abonné en fonction de son ID
// @ID delete-categorie-abonne
// @Produce json
// @Param id path string true "ID de la catégorie d'abonné à supprimer"
// @Success 204 "Catégorie d'abonné supprimée avec succès"
// @Router /api/categorie_abonne/{id} [delete]
func DeleteCategorieAbonneHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	vars := mux.Vars(r)
	idStr := vars["id"]

	err := db.DeleteCategorieAbonne(t, idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteTypeOuvrageHandler
// @Summary Supprimer un type d'ouvrage par ID
// @Description Supprime un type d'ouvrage en fonction de son ID
// @ID delete-type-ouvrage
// @Produce json
// @Param id path string true "ID du type d'ouvrage à supprimer"
// @Success 204 "Type d'ouvrage supprimé avec succès"
// @Router /api/type_ouvrage/{id} [delete]
func DeleteTypeOuvrageHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	vars := mux.Vars(r)
	idStr := vars["id"]

	err := db.DeleteTypeOuvrage(t, idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteAuteurHandler
// @Summary Supprimer un auteur par ID
// @Description Supprime un auteur en fonction de son ID
// @ID delete-auteur
// @Produce json
// @Param id path int true "ID de l'auteur à supprimer"
// @Success 204 "Auteur supprimé avec succès"
// @Router /api/auteur/{id} [delete]
func DeleteAuteurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.DeleteAuteur(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteEditeurHandler
// @Summary Supprimer un éditeur par ID
// @Description Supprime un éditeur en fonction de son ID
// @ID delete-editeur
// @Produce json
// @Param id path int true "ID de l'éditeur à supprimer"
// @Success 204 "Éditeur supprimé avec succès"
// @Router /api/editeur/{id} [delete]
func DeleteEditeurHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.DeleteEditeur(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteOuvrageHandler
// @Summary Supprimer un ouvrage par ID
// @Description Supprime un ouvrage en fonction de son ID
// @ID delete-ouvrage
// @Produce json
// @Param id path int true "ID de l'ouvrage à supprimer"
// @Success 204 "Ouvrage supprimé avec succès"
// @Router /api/ouvrage/{id} [delete]
func DeleteOuvrageHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.DeleteOuvrage(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteEcritHandler
// @Summary Supprimer un écrit par ID
// @Description Supprime un écrit en fonction de son ID
// @ID delete-ecrit
// @Produce json
// @Param id path int true "ID de l'écrit à supprimer"
// @Success 204 "Écrit supprimé avec succès"
// @Router /api/ecrit/{id} [delete]
func DeleteEcritHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.DeleteEcrit(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteEmpruntHandler
// @Summary Supprimer un emprunt par ID
// @Description Supprime un emprunt en fonction de son ID
// @ID delete-emprunt
// @Produce json
// @Param id path int true "ID de l'emprunt à supprimer"
// @Success 204 "Emprunt supprimé avec succès"
// @Router /api/emprunt/{id} [delete]
func DeleteEmpruntHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.DeleteEmprunter(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
