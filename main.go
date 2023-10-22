package main

import (
	"biblioV2/db"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Récupérer l'en-tête "Authorization" de la requête
		authHeader := r.Header.Get("Authorization")

		// Vérifier si l'en-tête "Authorization" est vide
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Diviser l'en-tête en deux parties : le type et le jeton
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		// Vérifier le type d'authentification (dans ce cas, nous supposons "Bearer")
		if parts[0] != "Bearer" {
			http.Error(w, "Unsupported Authorization type", http.StatusUnauthorized)
			return
		}

		// Récupérer le jeton d'authentification
		token := parts[1]

		// Vérifier le jeton (vous devrez définir votre propre logique pour vérifier le jeton)
		valid, err := isValidToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if !valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Si le jeton est valide, passez la demande au gestionnaire suivant
		next.ServeHTTP(w, r)
	})
}

func isValidToken(token string) (bool, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return false, err
	}

	return token == os.Getenv("TOKEN"), nil
}

func main() {
	apiRouter := mux.NewRouter()
	//apiRouter.Use(AuthMiddleware)

	t, err := db.OpenDBConnection()
	if err != nil {
		fmt.Errorf("error while opening database connection: %v", err)
	}
	defer t.Close()

	// GET
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

	// POST
	apiRouter.HandleFunc("/api/abonne", func(w http.ResponseWriter, r *http.Request) { PostAbonneHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/categorie_abonne", func(w http.ResponseWriter, r *http.Request) { PostCategorieAbonneHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/type_ouvrage", func(w http.ResponseWriter, r *http.Request) { PostTypeOuvrageHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/auteur", func(w http.ResponseWriter, r *http.Request) { PostAuteurHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/editeur", func(w http.ResponseWriter, r *http.Request) { PostEditeurHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/ouvrage", func(w http.ResponseWriter, r *http.Request) { PostOuvrageHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/ecrit", func(w http.ResponseWriter, r *http.Request) { PostEcritHandler(w, r, t) }).Methods("POST")
	apiRouter.HandleFunc("/api/emprunt", func(w http.ResponseWriter, r *http.Request) { PostEmpruntHandler(w, r, t) }).Methods("POST")

	// PUT
	apiRouter.HandleFunc("/api/abonne/{id}", func(w http.ResponseWriter, r *http.Request) { PutAbonneHandler(w, r, t) }).Methods("PUT")
	//router.HandleFunc("/categorie_abonne/{id}", PutCategorieAbonneHandler).Methods("PUT") - Pas de modification de la clé primaire
	//router.HandleFunc("/type_ouvrage/{id}", PutTypeOuvrageHandler).Methods("PUT") - Pas de modification de la clé primaire
	apiRouter.HandleFunc("/api/auteur/{id}", func(w http.ResponseWriter, r *http.Request) { PutAuteurHandler(w, r, t) }).Methods("PUT")
	apiRouter.HandleFunc("/api/editeur/{id}", func(w http.ResponseWriter, r *http.Request) { PutEditeurHandler(w, r, t) }).Methods("PUT")
	apiRouter.HandleFunc("/api/ouvrage/{id}", func(w http.ResponseWriter, r *http.Request) { PutOuvrageHandler(w, r, t) }).Methods("PUT")
	apiRouter.HandleFunc("/api/ecrit/{id}", func(w http.ResponseWriter, r *http.Request) { PutEcritHandler(w, r, t) }).Methods("PUT")
	apiRouter.HandleFunc("/api/emprunt/{id}", func(w http.ResponseWriter, r *http.Request) { PutEmpruntHandler(w, r, t) }).Methods("PUT")

	// DELETE
	apiRouter.HandleFunc("/api/abonne/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteAbonneHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/categorie_abonne/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteCategorieAbonneHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/type_ouvrage/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteTypeOuvrageHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/auteur/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteAuteurHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/editeur/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteEditeurHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/ouvrage/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteOuvrageHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/ecrit/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteEcritHandler(w, r, t) }).Methods("DELETE")
	apiRouter.HandleFunc("/api/emprunt/{id}", func(w http.ResponseWriter, r *http.Request) { DeleteEmpruntHandler(w, r, t) }).Methods("DELETE")

	//special case
	apiRouter.HandleFunc("/api/retour/{id}", func(w http.ResponseWriter, r *http.Request) { RetourEmprunterHandler(w, r, t) }).Methods("PUT")

	apiRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", apiRouter))

}
