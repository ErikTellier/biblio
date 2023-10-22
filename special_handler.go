package main

import (
	"biblioV2/db"
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// RetourEmprunterHandler
// @Summary Marquer un emprunt comme retourné
// @Description Marque un emprunt comme retourné en utilisant l'ID de l'emprunt fourni dans la requête DELETE
// @ID retour-emprunt
// @Produce json
// @Param id path int true "ID de l'emprunt à marquer comme retourné"
// @Success 200 "Emprunt marqué comme retourné avec succès"
// @Failure 500 "Échec de la mise à jour de l'état de l'emprunt"
// @Router /api/emprunt/retour/{id} [delete]
func RetourEmprunterHandler(w http.ResponseWriter, r *http.Request, t *sql.DB) {

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.RetourEmprunter(t, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
