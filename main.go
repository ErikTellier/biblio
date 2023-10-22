package main

import (
	"biblioV2/db"
	_ "biblioV2/docs"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

func main() {
	apiRouter := mux.NewRouter()

	t, err := db.OpenDBConnection()
	if err != nil {
		fmt.Errorf("error while opening database connection: %v", err)
	}
	defer t.Close()

	apiRouter.HandleFunc("/api/abonne", func(w http.ResponseWriter, r *http.Request) { GetAbonneHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/categorie_abonne", func(w http.ResponseWriter, r *http.Request) { GetCategorieAbonneHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/type_ouvrage", func(w http.ResponseWriter, r *http.Request) { GetTypeOuvrageHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/auteur", func(w http.ResponseWriter, r *http.Request) { GetAuteurHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/editeur", func(w http.ResponseWriter, r *http.Request) { GetEditeurHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/ouvrage", func(w http.ResponseWriter, r *http.Request) { GetOuvrageHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/ecrit", func(w http.ResponseWriter, r *http.Request) { GetEcritHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/emprunt", func(w http.ResponseWriter, r *http.Request) { GetEmprunterHandler(w, r, t) }).Methods("GET")

	apiRouter.HandleFunc("/api/abonne/{id}", func(w http.ResponseWriter, r *http.Request) { GetAbonneByIdHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/categorie_abonne/{id}", func(w http.ResponseWriter, r *http.Request) { GetCategorieAbonneByIdHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/type_ouvrage/{id}", func(w http.ResponseWriter, r *http.Request) { GetTypeOuvrageByIdHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/auteur/{id}", func(w http.ResponseWriter, r *http.Request) { GetAuteurByIdHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/editeur/{id}", func(w http.ResponseWriter, r *http.Request) { GetEditeurByIdHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/ouvrage/{id}", func(w http.ResponseWriter, r *http.Request) { GetOuvrageByIdHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/ecrit/{id}", func(w http.ResponseWriter, r *http.Request) { GetEcritByIdHandler(w, r, t) }).Methods("GET")
	apiRouter.HandleFunc("/api/emprunt/{id}", func(w http.ResponseWriter, r *http.Request) { GetEmpruntByIdHandler(w, r, t) }).Methods("GET")

	apiRouter.HandleFunc("/api/abonne", func(w http.ResponseWriter, r *http.Request) { PostAbonneHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/categorie_abonne", func(w http.ResponseWriter, r *http.Request) { PostCategorieAbonneHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/type_ouvrage", func(w http.ResponseWriter, r *http.Request) { PostTypeOuvrageHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/auteur", func(w http.ResponseWriter, r *http.Request) { PostAuteurHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/editeur", func(w http.ResponseWriter, r *http.Request) { PostEditeurHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/ouvrage", func(w http.ResponseWriter, r *http.Request) { PostOuvrageHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/ecrit", func(w http.ResponseWriter, r *http.Request) { PostEcritHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/emprunt", func(w http.ResponseWriter, r *http.Request) { PostEmpruntHandler(w, r, t) }).Methods("POST")

	apiRouter.HandleFunc("/api/abonne/{id}", func(w http.ResponseWriter, r *http.Request) { PutAbonneHandler(w, r, t) }).Methods("PUT")
	apiRouter.HandleFunc("/api/auteur/{id}", func(w http.ResponseWriter, r *http.Request) { PutAuteurHandler(w, r, t) }).Methods("PUT")
	apiRouter.HandleFunc("/api/editeur/{id}", func(w http.ResponseWriter, r *http.Request) { PutEditeurHandler(w, r, t) }).Methods("PUT")
	apiRouter.HandleFunc("/api/ouvrage/{id}", func(w http.ResponseWriter, r *http.Request) { PutOuvrageHandler(w, r, t) }).Methods("PUT")
	apiRouter.HandleFunc("/api/ecrit/{id}", func(w http.ResponseWriter, r *http.Request) { PutEcritHandler(w, r, t) }).Methods("PUT")
	apiRouter.HandleFunc("/api/emprunt/{id}", func(w http.ResponseWriter, r *http.Request) { PutEmpruntHandler(w, r, t) }).Methods("PUT")

	apiRouter.HandleFunc("/api/abonne/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteAbonneHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/categorie_abonne/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteCategorieAbonneHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/type_ouvrage/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteTypeOuvrageHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/auteur/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteAuteurHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/editeur/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteEditeurHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/ouvrage/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteOuvrageHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/ecrit/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteEcritHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/emprunt/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteEmpruntHandler(w, r, t) }).Methods("DELETE")

	apiRouter.HandleFunc("/api/retour/{id}", func(w http.ResponseWriter, r *http.Request) { RetourEmprunterHandler(w, r, t) }).Methods("PUT")

	apiRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	apiRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", apiRouter))

}
