package main

import (
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

	// GET
	apiRouter.HandleFunc("/api/abonne", GetAbonneHandler).Methods("GET")
	apiRouter.HandleFunc("/api/categorie_abonne", GetCategorieAbonneHandler).Methods("GET")
	apiRouter.HandleFunc("/api/type_ouvrage", GetTypeOuvrageHandler).Methods("GET")
	apiRouter.HandleFunc("/api/auteur", GetAuteurHandler).Methods("GET")
	apiRouter.HandleFunc("/api/editeur", GetEditeurHandler).Methods("GET")
	apiRouter.HandleFunc("/api/ouvrage", GetOuvrageHandler).Methods("GET")
	apiRouter.HandleFunc("/api/ecrit", GetEcritHandler).Methods("GET")
	apiRouter.HandleFunc("/api/emprunt", GetEmprunterHandler).Methods("GET")

	apiRouter.HandleFunc("/api/abonne/{id}", GetAbonneByIdHandler).Methods("GET")
	apiRouter.HandleFunc("/api/categorie_abonne/{id}", GetCategorieAbonneByIdHandler).Methods("GET")
	apiRouter.HandleFunc("/api/type_ouvrage/{id}", GetTypeOuvrageByIdHandler).Methods("GET")
	apiRouter.HandleFunc("/api/auteur/{id}", GetAuteurByIdHandler).Methods("GET")
	apiRouter.HandleFunc("/api/editeur/{id}", GetEditeurByIdHandler).Methods("GET")
	apiRouter.HandleFunc("/api/ouvrage/{id}", GetOuvrageByIdHandler).Methods("GET")
	apiRouter.HandleFunc("/api/ecrit/{id}", GetEcritByIdHandler).Methods("GET")
	apiRouter.HandleFunc("/api/emprunt/{id}", GetEmpruntByIdHandler).Methods("GET")

	// POST
	apiRouter.HandleFunc("/api/abonne", PostAbonneHandler).Methods("POST")
	apiRouter.HandleFunc("/api/categorie_abonne", PostCategorieAbonneHandler).Methods("POST")
	apiRouter.HandleFunc("/api/type_ouvrage", PostTypeOuvrageHandler).Methods("POST")
	apiRouter.HandleFunc("/api/auteur", PostAuteurHandler).Methods("POST")
	apiRouter.HandleFunc("/api/editeur", PostEditeurHandler).Methods("POST")
	apiRouter.HandleFunc("/api/ouvrage", PostOuvrageHandler).Methods("POST")
	apiRouter.HandleFunc("/api/ecrit", PostEcritHandler).Methods("POST")
	apiRouter.HandleFunc("/api/emprunt", PostEmpruntHandler).Methods("POST")

	// PUT
	apiRouter.HandleFunc("/api/abonne/{id}", PutAbonneHandler).Methods("PUT")
	//router.HandleFunc("/categorie_abonne/{id}", PutCategorieAbonneHandler).Methods("PUT") - Pas de modification de la clé primaire
	//router.HandleFunc("/type_ouvrage/{id}", PutTypeOuvrageHandler).Methods("PUT") - Pas de modification de la clé primaire
	apiRouter.HandleFunc("/api/auteur/{id}", PutAuteurHandler).Methods("PUT")
	apiRouter.HandleFunc("/api/editeur/{id}", PutEditeurHandler).Methods("PUT")
	apiRouter.HandleFunc("/api/ouvrage/{id}", PutOuvrageHandler).Methods("PUT")
	apiRouter.HandleFunc("/api/ecrit/{id}", PutEcritHandler).Methods("PUT")
	apiRouter.HandleFunc("/api/emprunt/{id}", PutEmpruntHandler).Methods("PUT")

	// DELETE
	apiRouter.HandleFunc("/api/abonne/{id}", DeleteAbonneHandler).Methods("DELETE")
	apiRouter.HandleFunc("/api/categorie_abonne/{id}", DeleteCategorieAbonneHandler).Methods("DELETE")
	apiRouter.HandleFunc("/api/type_ouvrage/{id}", DeleteTypeOuvrageHandler).Methods("DELETE")
	apiRouter.HandleFunc("/api/auteur/{id}", DeleteAuteurHandler).Methods("DELETE")
	apiRouter.HandleFunc("/api/editeur/{id}", DeleteEditeurHandler).Methods("DELETE")
	apiRouter.HandleFunc("/api/ouvrage/{id}", DeleteOuvrageHandler).Methods("DELETE")
	apiRouter.HandleFunc("/api/ecrit/{id}", DeleteEcritHandler).Methods("DELETE")
	apiRouter.HandleFunc("/api/emprunt/{id}", DeleteEmpruntHandler).Methods("DELETE")

	//special case
	apiRouter.HandleFunc("/api/retour/{id}", RetourEmprunterHandler).Methods("PUT")

	apiRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", apiRouter))

}
