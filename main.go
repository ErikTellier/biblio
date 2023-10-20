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

	router.HandleFunc("/abonne/{id}", GetAbonneByIdHandler).Methods("GET")
	router.HandleFunc("/categorie_abonne/{id}", GetCategorieAbonneByIdHandler).Methods("GET")
	router.HandleFunc("/type_ouvrage/{id}", GetTypeOuvrageByIdHandler).Methods("GET")
	router.HandleFunc("/auteur/{id}", GetAuteurByIdHandler).Methods("GET")
	router.HandleFunc("/editeur/{id}", GetEditeurByIdHandler).Methods("GET")
	router.HandleFunc("/ouvrage/{id}", GetOuvrageByIdHandler).Methods("GET")
	router.HandleFunc("/ecrit/{id}", GetEcritByIdHandler).Methods("GET")
	router.HandleFunc("/emprunt/{id}", GetEmpruntByIdHandler).Methods("GET")

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
	router.HandleFunc("/abonne/{id}", PutAbonneHandler).Methods("PUT")
	//router.HandleFunc("/categorie_abonne/{id}", PutCategorieAbonneHandler).Methods("PUT") - Pas de modification de la clé primaire
	//router.HandleFunc("/type_ouvrage/{id}", PutTypeOuvrageHandler).Methods("PUT") - Pas de modification de la clé primaire
	router.HandleFunc("/auteur/{id}", PutAuteurHandler).Methods("PUT")
	router.HandleFunc("/editeur/{id}", PutEditeurHandler).Methods("PUT")
	router.HandleFunc("/ouvrage/{id}", PutOuvrageHandler).Methods("PUT")
	router.HandleFunc("/ecrit/{id}", PutEcritHandler).Methods("PUT")
	router.HandleFunc("/emprunt/{id}", PutEmpruntHandler).Methods("PUT")

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
