package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// GET
	router.HandleFunc("/abonne", GetAbonneHandler).Methods("GET")
	router.HandleFunc("/categorie_abonne", GetCategorieAbonneHandler).Methods("GET")
	router.HandleFunc("/type_ouvrage", GetTypeOuvrageHandler).Methods("GET")
	router.HandleFunc("/auteur", GetAuteurHandler).Methods("GET")
	router.HandleFunc("/editeur", GetEditeurHandler).Methods("GET")
	router.HandleFunc("/ouvrage", GetOuvrageHandler).Methods("GET")
	router.HandleFunc("/ecrit", GetEcritHandler).Methods("GET")
	router.HandleFunc("/emprunt", GetEmprunterHandler).Methods("GET")

	// POST
	router.HandleFunc("/abonne", PostAbonneHandler).Methods("POST")
	router.HandleFunc("/categorie_abonne", PostCategorieAbonneHandler).Methods("POST")
	router.HandleFunc("/type_ouvrage", PostTypeOuvrageHandler).Methods("POST")
	router.HandleFunc("/auteur", PostAuteurHandler).Methods("POST")
	router.HandleFunc("/editeur", PostEditeurHandler).Methods("POST")
	router.HandleFunc("/ouvrage", PostOuvrageHandler).Methods("POST")
	router.HandleFunc("/ecrit", PostEcritHandler).Methods("POST")
	router.HandleFunc("/emprunt", PostEmpruntHandler).Methods("POST")

	// PUT

	// DELETE
	router.HandleFunc("/abonne/{id}", DeleteAbonneHandler).Methods("DELETE")
	router.HandleFunc("/categorie_abonne/{id}", DeleteCategorieAbonneHandler).Methods("DELETE")
	router.HandleFunc("/type_ouvrage/{id}", DeleteTypeOuvrageHandler).Methods("DELETE")
	router.HandleFunc("/auteur/{id}", DeleteAuteurHandler).Methods("DELETE")
	router.HandleFunc("/editeur/{id}", DeleteEditeurHandler).Methods("DELETE")
	router.HandleFunc("/ouvrage/{id}", DeleteOuvrageHandler).Methods("DELETE")
	router.HandleFunc("/ecrit/{id}", DeleteEcritHandler).Methods("DELETE")
	router.HandleFunc("/emprunt/{id}", DeleteEmpruntHandler).Methods("DELETE")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
